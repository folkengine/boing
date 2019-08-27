package cmd

import (
	"github.com/folkengine/boing"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Makes a submarine sonar ping sound",
	Run: func(cmd *cobra.Command, args []string) {
		boing.Ping()
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
