package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "phpipamsync",
	Short: "phpipamsync is a CLI tool to synchronize phpIPAM and pi-hole",
	Long:  "phpipamsync pulls subnet from phpIPAM and generate either Cisco DHCP config or custom.file for pi-hole",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
