package cmd

import (
	"fmt"
	"instago/pkg"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs into an instagram accout with user credentials",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		fmt.Println("config file path", cfgFile)
		pkg.GetUserCredentials(cfgFile)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
