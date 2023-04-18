package cmd

import (
	"fmt"
	"os"

	apiclient "github.com/ivankuchin/sdwanctl/internal/api_client"
	config_reader "github.com/ivankuchin/sdwanctl/internal/config-reader"
	"github.com/spf13/cobra"
)

var describeTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Get template info",
	Long:  "Get template info",
}

var describeTemplateInputCmd = &cobra.Command{
	Use:   "input",
	Short: "Get feature template input",
	Long:  "Get feature template input",
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

		templ_input_list, err := apiclient.DescribeTemplateInputByDevID(config_reader.Cfg, device_id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if len(templ_input_list.TemplateInput) == 0 {
			fmt.Println("No devices found")
			os.Exit(1)
		}

		system_device := templ_input_list.TemplateInput[0]
		describeTemplateInput(system_device)
	},
}

func init() {
	describeCmd.AddCommand(describeTemplateCmd)
	describeTemplateCmd.AddCommand(describeTemplateInputCmd)
}

func describeTemplateInput(templ_input apiclient.TemplateInput) {
	for k, v := range templ_input.Entry {
		fmt.Printf("%s: %s\n", k, v)
	}
}
