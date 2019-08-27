package cmd

import (
	"github.com/folkengine/boing"

	"github.com/spf13/cobra"
)

// klaxonCmd represents the klaxon command
var klaxonCmd = &cobra.Command{
	Use:   "klaxon",
	Short: "DANGER! DANGER",
	Run: func(cmd *cobra.Command, args []string) {
		boing.Klaxon()
	},
}

func init() {
	rootCmd.AddCommand(klaxonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// klaxonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// klaxonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
