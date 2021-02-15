package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs into an instagram accout with user credentials",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
