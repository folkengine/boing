package cmd

import (
	"github.com/folkengine/boing"

	"github.com/spf13/cobra"
)

// crowCmd represents the rooster command
var crowCmd = &cobra.Command{
	Use:   "bock",
	Short: "Bock Bock Bitches!make" +
		"",
	Run: func(cmd *cobra.Command, args []string) {
		boing.Crow()
	},
}

func init() {
	rootCmd.AddCommand(crowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
