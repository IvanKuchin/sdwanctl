package cmd

import (
	"fmt"
	"os"

	apiclient "github.com/ivankuchin/sdwanctl/internal/api_client"
	config_reader "github.com/ivankuchin/sdwanctl/internal/config-reader"
	"github.com/spf13/cobra"
)

var describeSystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Get system info",
	Long:  "Get system info",
}

var describeSystemDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Get detailed system device info",
	Long:  "Get detailed system device info",
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

		dev, err := apiclient.DescribeSystemDeviceByDevID(config_reader.Cfg, device_id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if len(dev.Devices) == 0 {
			fmt.Println("No devices found")
			os.Exit(1)
		}

		system_device := dev.Devices[0]
		describeSystemDevice(system_device)
	},
}

func init() {
	describeCmd.AddCommand(describeSystemCmd)
	describeSystemCmd.AddCommand(describeSystemDeviceCmd)
}

func describeSystemDevice(dev apiclient.SystemDevice) {
	fmt.Printf("DeviceType: %v\n", dev.DeviceType)
	fmt.Printf("SerialNumber: %v\n", dev.SerialNumber)
	fmt.Printf("NcsDeviceName: %v\n", dev.NcsDeviceName)
	fmt.Printf("ConfigStatusMessage: %v\n", dev.ConfigStatusMessage)
	fmt.Printf("TemplateApplyLog: %v\n", dev.TemplateApplyLog)
	fmt.Printf("UUID: %v\n", dev.UUID)
	fmt.Printf("ManagementSystemIP: %v\n", dev.ManagementSystemIP)
	fmt.Printf("TemplateStatus: %v\n", dev.TemplateStatus)
	fmt.Printf("ChasisNumber: %v\n", dev.ChasisNumber)
	fmt.Printf("ConfigStatusMessageDetails: %v\n", dev.ConfigStatusMessageDetails)
	fmt.Printf("ConfigOperationMode: %v\n", dev.ConfigOperationMode)
	fmt.Printf("DeviceModel: %v\n", dev.DeviceModel)
	fmt.Printf("DeviceState: %v\n", dev.DeviceState)
	fmt.Printf("Validity: %v\n", dev.Validity)
	fmt.Printf("PlatformFamily: %v\n", dev.PlatformFamily)
	fmt.Printf("VedgeCertificateState: %v\n", dev.VedgeCertificateState)
	fmt.Printf("RootCertHash: %v\n", dev.RootCertHash)
	fmt.Printf("DeviceIP: %v\n", dev.DeviceIP)
	fmt.Printf("Personality: %v\n", dev.Personality)
	fmt.Printf("UploadSource: %v\n", dev.UploadSource)
	fmt.Printf("SubjectSerialNumber: %v\n", dev.SubjectSerialNumber)
	fmt.Printf("LocalSystemIP: %v\n", dev.LocalSystemIP)
	fmt.Printf("SystemIP: %v\n", dev.SystemIP)
	fmt.Printf("ModelSku: %v\n", dev.ModelSku)
	fmt.Printf("SiteID: %v\n", dev.SiteID)
	fmt.Printf("HostName: %v\n", dev.HostName)
	fmt.Printf("Version: %v\n", dev.Version)
	fmt.Printf("Vbond: %v\n", dev.Vbond)
	fmt.Printf("VmanageConnectionState: %v\n", dev.VmanageConnectionState)
	fmt.Printf("Lastupdated: %v\n", dev.Lastupdated)
	fmt.Printf("Reachability: %v\n", dev.Reachability)
	fmt.Printf("UptimeDate: %v\n", dev.UptimeDate)
	fmt.Printf("DefaultVersion: %v\n", dev.DefaultVersion)
	fmt.Printf("AvailableVersions: %v\n", dev.AvailableVersions)
	fmt.Printf("Template: %v\n", dev.Template)
	fmt.Printf("TemplateID: %v\n", dev.TemplateID)
	fmt.Printf("LifeCycleRequired: %v\n", dev.LifeCycleRequired)
	fmt.Printf("ExpirationDate: %v\n", dev.ExpirationDate)
	fmt.Printf("HardwareCertSerialNumber: %v\n", dev.HardwareCertSerialNumber)
	fmt.Printf("DraftMode: %v\n", dev.DraftMode)

}
