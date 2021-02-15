package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ghostsCmd represents the ghosts command
var ghostsCmd = &cobra.Command{
	Use:   "ghosts",
	Short: "Checks which users don't follow you back",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: from accounts that don't follow you back, how many are verified & how many are not
		fmt.Println("ghosts called")
	},
}

func init() {
	trackCmd.AddCommand(ghostsCmd)
}
