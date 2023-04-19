package cmd

import (
	"fmt"
	"os"

	apiclient "github.com/ivankuchin/sdwanctl/internal/api_client"
	config_reader "github.com/ivankuchin/sdwanctl/internal/config-reader"
	"github.com/spf13/cobra"
)

var repostCmd = &cobra.Command{
	Use:   "repost",
	Short: "Repost to a device w/o changes",
	Long:  "Repost to a device w/o changes",
}

var repostTemplateInputCmd = &cobra.Command{
	Use:   "template",
	Short: "Repost to a device w/o changes",
	Long:  "Repost to a device w/o changes",
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

		template_input := templ_input_list.TemplateInput[0]

		err = apiclient.RepostTemplate(config_reader.Cfg, template_input)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println("Repost to deviceID " + device_id + " is successful")
	},
}

func init() {
	rootCmd.AddCommand(repostCmd)
	repostCmd.AddCommand(repostTemplateInputCmd)
}
