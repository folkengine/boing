package cmd

import (
	"github.com/folkengine/boing"
	"github.com/spf13/cobra"
)

// toneCmd represents the tone command
var toneCmd = &cobra.Command{
	Use:   "tone",
	Short: "Plays a clear tone",
	Run: func(cmd *cobra.Command, args []string) {
		boing.Tone()
	},
}

func init() {
	rootCmd.AddCommand(toneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
