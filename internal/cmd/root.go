package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sdwanctl",
	Short: "sdwanctl is a CLI tool for managing Cisco SD-WAN",
	Long:  "sdwanctl is a CLI tool for managing Cisco SD-WAN",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
