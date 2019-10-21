package cmd

import (
	"github.com/folkengine/boing"

	"github.com/spf13/cobra"
)

// bockCmd represents the bock command
var bockCmd = &cobra.Command{
	Use:   "crow",
	Short: "Good morning!!!" +
		"",
	Run: func(cmd *cobra.Command, args []string) {
		boing.Crow()
	},
}

func init() {
	rootCmd.AddCommand(bockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
