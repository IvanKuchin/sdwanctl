package cmd

import (
	"fmt"
	"log"
	"os"

	apiclient "github.com/ivankuchin/sdwanctl/internal/api_client"
	config_reader "github.com/ivankuchin/sdwanctl/internal/config-reader"
	"github.com/spf13/cobra"
)

var getConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get device config",
	Long:  "Get device config",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please specify a single deviceID")
			os.Exit(1)
		}
		device_id := args[0]

		err := apiclient.Login(config_reader.Cfg)
		if err != nil {
			os.Exit(1)
		}
		defer apiclient.Logout(config_reader.Cfg)

		device_config, err := apiclient.GetDeviceConfig(config_reader.Cfg, device_id)
		if err != nil {
			log.Fatal(err.Error())
		}

		printDeviceConfig(device_id, device_config)
	},
}

func init() {
	getCmd.AddCommand(getConfigCmd)
}

func printDeviceConfig(device_id, device_config string) {
	fmt.Printf("DEvice ID: %s\n--------------------\n%s", device_id, device_config)
}
