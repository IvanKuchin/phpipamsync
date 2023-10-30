package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ivankuchin/phpipamsync/internal/api_client"
	"github.com/ivankuchin/phpipamsync/internal/config_reader"
	"github.com/spf13/cobra"
)

func getSubnetID(client api_client.Authenticator, cfg *config_reader.Config) (int, error) {
	subnet_id := 0
	subnet_json, err := client.Call("GET", cfg.Ipam_site_url+"/api/"+cfg.Ipam_app_id+"/subnets/cidr/"+cfg.Ipam_subnet, "")
	if err != nil {
		return subnet_id, err
	}

	subnets := IpamSubnet_list{}
	err = json.Unmarshal(subnet_json, &subnets)
	if err != nil {
		log.Printf("Error unmarshalling json: %s", err.Error())
		return subnet_id, err
	}

	subnet_id, err = strconv.Atoi(subnets.Subnets[0].Id)
	if err != nil {
		log.Printf("Error converting subnet id to int: %s", err.Error())
		return subnet_id, err
	}

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

func getPiHoleCustomOutput(addresses IPAddresses, cfg *config_reader.Config) string {
	output := ""

	for _, address := range addresses.IPAddresses {
		switch address.Tag {
		case "2": // In Use
			if address.Hostname == "" {
				log.Printf("WARNING: Skipping %s because hostname is empty", address.IP)
			} else if address.IP == "" {
				log.Printf("WARNING: Skipping %s because IP is empty", address.Hostname)
			} else {
				output += address.IP + " " + address.Hostname + "." + cfg.Domain + "\n"
			}
		case "3": // Reserved
		case "4": // DHCP pool
		default:
			log.Printf("WARNING: Skipping %s because tag is %s", address.IP, address.Tag)
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

		auth := new(api_client.AuthAppCode)
		err := auth.Login(config_reader.Cfg)
		if err != nil {
			return err
		}
		defer auth.Logout()

		subnet_id, err := getSubnetID(auth, config_reader.Cfg)
		if err != nil {
			return err
		}

		addresses, err := getIPAddressesBySubnetID(subnet_id, auth, config_reader.Cfg)
		if err != nil {
			return err
		}

		fmt.Println(addresses)

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

		subnet_id, err := getSubnetID(auth, config_reader.Cfg)
		if err != nil {
			return err
		}

		addresses, err := getIPAddressesBySubnetID(subnet_id, auth, config_reader.Cfg)
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
