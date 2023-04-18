package cmd

import (
	"fmt"
	"log"
	"os"

	apiclient "github.com/ivankuchin/sdwanctl/internal/api_client"
	config_reader "github.com/ivankuchin/sdwanctl/internal/config-reader"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Get information about single entity",
	Long:  "Get information about single entity",
}

var describeDevicesCmd = &cobra.Command{
	Use:   "device",
	Short: "Get device info",
	Long:  "Get device info",
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

		dev, err := apiclient.DescribeDeviceByDevID(config_reader.Cfg, device_id)
		if err != nil {
			log.Fatal(err.Error())
		}

		describeDevice(dev)
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
	describeCmd.AddCommand(describeDevicesCmd)
}

func describeDevice(dev apiclient.Device) {
	fmt.Printf("DeviceID: %v\n", dev.DeviceID)
	fmt.Printf("SystemIP: %v\n", dev.SystemIP)
	fmt.Printf("HostName: %v\n", dev.HostName)
	fmt.Printf("Reachability: %v\n", dev.Reachability)
	fmt.Printf("Status: %v\n", dev.Status)
	fmt.Printf("UUID: %v\n", dev.UUID)
	fmt.Printf("Personality: %v\n", dev.Personality)
	fmt.Printf("DeviceType: %v\n", dev.DeviceType)
	fmt.Printf("Timezone: %v\n", dev.Timezone)
	fmt.Printf("DeviceGroups: %v\n", dev.DeviceGroups)
	fmt.Printf("Lastupdated: %v\n", dev.Lastupdated)
	fmt.Printf("DomainID: %v\n", dev.DomainID)
	fmt.Printf("BoardSerial: %v\n", dev.BoardSerial)
	fmt.Printf("CertificateValidity: %v\n", dev.CertificateValidity)
	fmt.Printf("MaxControllers: %v\n", dev.MaxControllers)
	fmt.Printf("ControlConnections: %v\n", dev.ControlConnections)
	fmt.Printf("DeviceModel: %v\n", dev.DeviceModel)
	fmt.Printf("Version: %v\n", dev.Version)
	fmt.Printf("ConnectedVManages: %v\n", dev.ConnectedVManages)
	fmt.Printf("SiteID: %v\n", dev.SiteID)
	fmt.Printf("Latitude: %v\n", dev.Latitude)
	fmt.Printf("Longitude: %v\n", dev.Longitude)
	fmt.Printf("IsDeviceGeoData: %v\n", dev.IsDeviceGeoData)
	fmt.Printf("Platform: %v\n", dev.Platform)
	fmt.Printf("UptimeDate: %v\n", dev.UptimeDate)
	fmt.Printf("StatusOrder: %v\n", dev.StatusOrder)
	fmt.Printf("DeviceOs: %v\n", dev.DeviceOs)
	fmt.Printf("Validity: %v\n", dev.Validity)
	fmt.Printf("State: %v\n", dev.State)
	fmt.Printf("StateDescription: %v\n", dev.StateDescription)
	fmt.Printf("ModelSku: %v\n", dev.ModelSku)
	fmt.Printf("LocalSystemIP: %v\n", dev.LocalSystemIP)
	fmt.Printf("TotalCPUCount: %v\n", dev.TotalCPUCount)
	fmt.Printf("TestbedMode: %v\n", dev.TestbedMode)
	fmt.Printf("LayoutLevel: %v\n", dev.LayoutLevel)
	fmt.Printf("OmpPeers: %v\n", dev.OmpPeers)
	fmt.Printf("LinuxCPUCount: %v\n", dev.LinuxCPUCount)
	fmt.Printf("BfdSessionsUp: %v\n", dev.BfdSessionsUp)
	fmt.Printf("BfdSessions: %v\n", dev.BfdSessions)
}
