package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information from phpIPAM",
	Long:  "Get information from phpIPAM",
}

var getDevicesCmd = &cobra.Command{
	Use:   "cisco-dhcp",
	Short: "Get DHCP config for Cisco IOS devices",
	Long:  "Get DHCP config for Cisco IOS devices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add code here to get DHCP config for Cisco IOS devices")
		// err := apiclient.Login(config_reader.Cfg)
		// 	if err != nil {
		// 		os.Exit(1)
		// 	}
		// 	defer apiclient.Logout(config_reader.Cfg)

		// 	devices, err := apiclient.GetDeviceList(config_reader.Cfg)
		// 	if err != nil {
		// 		log.Fatal(err.Error())
		// 	}

		// 	printDeviceList(devices)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getDevicesCmd)
}
