package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track a specific metric related to your account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("track called")
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
