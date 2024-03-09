package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/netip"
	"os"
	"strconv"
	"strings"

	"github.com/ivankuchin/phpipamsync/internal/api_client"
	"github.com/ivankuchin/phpipamsync/internal/config_reader"
	"github.com/spf13/cobra"
)

func getSubnetID(client api_client.Authenticator, cfg *config_reader.Config, subnet string) (int, error) {
	subnet_id := 0
	subnet_json, err := client.Call("GET", cfg.Ipam_site_url+"/api/"+cfg.Ipam_app_id+"/subnets/cidr/"+subnet, "")
	if err != nil {
		return subnet_id, err
	}

	subnets := IpamSubnet_list{}
	err = json.Unmarshal(subnet_json, &subnets)
	if err != nil {
		log.Printf("Error unmarshalling json: %s", err.Error())
		return subnet_id, err
	}

	subnet_id = subnets.Subnets[0].Id

	return subnet_id, nil
}

func getIPAddressesBySubnetID(subnetID int, client api_client.Authenticator, cfg *config_reader.Config) (IPAddresses, error) {
	addresses := IPAddresses{}

	addresses_json, err := client.Call("GET", cfg.Ipam_site_url+"/api/"+cfg.Ipam_app_id+"/subnets/"+strconv.Itoa(subnetID)+"/addresses/", "")
	if err != nil {
		return addresses, err
	}

	err = json.Unmarshal(addresses_json, &addresses)
	if err != nil {
		log.Printf("Error unmarshalling json: %s", err.Error())
		return addresses, err
	}

	return addresses, nil
}

func getIPAddressesBelongsToSubnets(client api_client.Authenticator, cfg *config_reader.Config) (IPAddresses, error) {

	addresses := IPAddresses{}

	if len(cfg.Ipam_subnets) == 0 {
		return addresses, fmt.Errorf("no subnets defined in config file")
	}

	for _, subnet := range cfg.Ipam_subnets {

		subnet_id, err := getSubnetID(client, config_reader.Cfg, subnet)
		if err != nil {
			return addresses, err
		}

		addresses_temp, err := getIPAddressesBySubnetID(subnet_id, client, config_reader.Cfg)
		if err != nil {
			return addresses, err
		}

		addresses.IPAddresses = append(addresses.IPAddresses, addresses_temp.IPAddresses...)

	}
	return addresses, nil
}

func getPiHoleCustomOutput(addresses IPAddresses, cfg *config_reader.Config) string {
	output := ""

	for _, address := range addresses.IPAddresses {
		switch address.Tag {
		case 2: // In Use
			if address.Hostname == "" {
				log.Printf("WARNING: Skipping %s because hostname is empty", address.IP)
			} else if address.IP == "" {
				log.Printf("WARNING: Skipping %s because IP is empty", address.Hostname)
			} else {
				output += address.IP + " " + address.Hostname + "." + cfg.Domain + "\n"
			}
		case 3: // Reserved
		case 4: // DHCP pool
		default:
			log.Printf("WARNING: Skipping %s because tag is %v", address.IP, address.Tag)
		}
	}

	return output
}

func writeToPiHoleCustom(addresses string, cfg *config_reader.Config) error {

	// write string "addresses" to a file
	err := os.WriteFile(cfg.Pi_hole, []byte(addresses), 0644)
	if err != nil {
		log.Printf("ERROR writing to file: %s", err.Error())
		return err
	}

	return nil
}

func getSubnetMaskBySubnetID(subnetID string) (string, error) {
	mask := ""

	prefix, err := netip.ParsePrefix(subnetID)
	if err != nil {
		log.Printf("Error parsing subnet ID: %s", err.Error())
		return mask, err
	}

	bits := prefix.Bits()
	mask = net.IP(net.CIDRMask(bits, 32)).String()

	return mask, nil
}

func getGWIPBySubnetID(subnetID string) (string, error) {
	ip_addr := ""

	prefix, err := netip.ParsePrefix(subnetID)
	if err != nil {
		log.Printf("Error parsing subnet ID: %s", err.Error())
		return ip_addr, err
	}

	ip_addr = prefix.Addr().Next().String()

	return ip_addr, nil
}

func convertMacToClientIdentifier(mac string) string {
	cleaned_mac := strings.ReplaceAll(mac, ":", "")
	client_identifier := ""

	for i := 0; i < len(cleaned_mac); i++ {
		if (i-2)%4 == 0 {
			client_identifier += "."
		}
		client_identifier += string(cleaned_mac[i])
	}

	return "01" + client_identifier
}

func getCiscoDHCPOutputBySubnet(addresses IPAddresses, subnet string) (string, error) {
	output := ""

	mask, err := getSubnetMaskBySubnetID(subnet)
	if err != nil {
		return output, err
	}

	gw, err := getGWIPBySubnetID(subnet)
	if err != nil {
		return output, err
	}

	for _, address := range addresses.IPAddresses {
		switch address.Tag {
		case 2: // In Use
			if address.Hostname == "" {
				log.Printf("WARNING: Skipping %s because hostname is empty", address.IP)
			} else if address.IP == "" {
				log.Printf("WARNING: Skipping %s because IP is empty", address.Hostname)
			} else {
				output += "ip dhcp pool " + address.Hostname + "\n"
				output += " host " + address.IP + " " + mask + "\n"
				switch {
				case address.ClientIdentifier != "":
					output += " client-identifier " + address.ClientIdentifier + "\n"
				case address.Mac != "":
					output += " client-identifier " + convertMacToClientIdentifier(address.Mac) + "\n"
				}
				output += " default-router " + gw + "\n"
				output += "!\n"
			}
		case 3: // Reserved
		case 4: // DHCP pool
		default:
			log.Printf("WARNING: Skipping %s because tag is %v", address.IP, address.Tag)
		}
	}

	return output, nil
}

func getCiscoDHCPOutput(client api_client.Authenticator, cfg *config_reader.Config) (string, error) {
	output := ""

	if len(cfg.Ipam_subnets) == 0 {
		return "", fmt.Errorf("no subnets defined in config file")
	}

	for _, subnet := range cfg.Ipam_subnets {

		subnet_id, err := getSubnetID(client, config_reader.Cfg, subnet)
		if err != nil {
			return "", err
		}

		addresses, err := getIPAddressesBySubnetID(subnet_id, client, config_reader.Cfg)
		if err != nil {
			return "", err
		}

		temp_output, err := getCiscoDHCPOutputBySubnet(addresses, subnet)
		if err != nil {
			return "", err
		}

		output += temp_output
	}

	return output, nil
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information from phpIPAM",
	Long:  "Get information from phpIPAM",
}

var getCiscoDHCP = &cobra.Command{
	Use:   "cisco-dhcp",
	Short: "Get DHCP config for Cisco IOS devices",
	Long:  "Get DHCP config for Cisco IOS devices",
	RunE: func(cmd *cobra.Command, args []string) error {

		auth := api_client.AuthAppCode{}
		err := auth.Login(config_reader.Cfg)
		if err != nil {
			return err
		}
		defer auth.Logout()

		output, err := getCiscoDHCPOutput(&auth, config_reader.Cfg)
		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	},
}

var getPiHoleCustom = &cobra.Command{
	Use:   "pi-hole",
	Short: "Get pi-hole config for custom.devices",
	Long:  "Get pi-hole config for custom.devices",
	RunE: func(cmd *cobra.Command, args []string) error {

		auth := new(api_client.AuthAppCode)
		err := auth.Login(config_reader.Cfg)
		if err != nil {
			return err
		}
		defer auth.Logout()

		addresses, err := getIPAddressesBelongsToSubnets(auth, config_reader.Cfg)
		if err != nil {
			return err
		}

		output := getPiHoleCustomOutput(addresses, config_reader.Cfg)

		err = writeToPiHoleCustom(output, config_reader.Cfg)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getCiscoDHCP)
	getCmd.AddCommand(getPiHoleCustom)
}
