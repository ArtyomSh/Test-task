package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "TestTask",
	Short: "Application for obtaining rates",
	Long: `You can use this application to run server and get Rates from it.
	Supported commands: 
		-server: run server (default port 3001)
		-rate: parameters:
		*pairs - pairs whose prices you need to get`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
