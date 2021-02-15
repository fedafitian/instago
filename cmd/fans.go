package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fansCmd represents the fans command
var fansCmd = &cobra.Command{
	Use:   "fans",
	Short: "Checks which users follow you that you don't follow back",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fans called")
	},
}

func init() {
	trackCmd.AddCommand(fansCmd)
}
