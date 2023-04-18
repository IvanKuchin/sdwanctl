package cmd

import (
	"fmt"
	"log"
	"os"

	apiclient "github.com/ivankuchin/sdwanctl/internal/api_client"
	config_reader "github.com/ivankuchin/sdwanctl/internal/config-reader"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information from Cisco SD-WAN",
	Long:  "Get information from Cisco SD-WAN",
}

var getDevicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Get device list",
	Long:  "Get device list",
	Run: func(cmd *cobra.Command, args []string) {
		err := apiclient.Login(config_reader.Cfg)
		if err != nil {
			os.Exit(1)
		}
		defer apiclient.Logout(config_reader.Cfg)

		devices, err := apiclient.GetDeviceList(config_reader.Cfg)
		if err != nil {
			log.Fatal(err.Error())
		}

		printDeviceList(devices)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getDevicesCmd)
}

func printDeviceList(dl apiclient.DeviceList) {
	fmt.Printf("%-45s %-20s %-20s %-28s %-10s %-20s %-10s\n", "deviceID", "HostName", "SystemIP", "DeviceModel", "DeviceType", "Version", "State")
	for _, device := range dl.Devices {
		fmt.Printf("%-45s %-20s %-20s %-28s %-10s %-20s %-10s\n", device.DeviceID, device.HostName, device.SystemIP, device.DeviceModel, device.DeviceType, device.Version, device.State)
	}
}
