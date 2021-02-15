package pkg

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	"github.com/spf13/viper"
)

type Account struct {
	//TODO: add instagram golang api
	Username  string
	Password  string
	Auth      bool // successful login status
	Following map[string]bool
	Followers map[string]bool
	Ghosts    []string //users who dont follow back
}

func readCredentialsFromConfig() (Account, error) {

	if _, err := os.Stat("~/.instago.env"); os.IsNotExist(err){
		// if config does not exist/cannot read credentials from config
		// ask, would you like to create a config file?
		// y- print add config and then try again
		// n- get user & pass via user input
		// if user does not have config file/cannot login with credentials from config
		var username string
		fmt.Scanln("Username: ")
		fmt.Scanln(&username)
		if err != nil {
		}
		fmt.Println("Password: ")
		password, err := terminal.ReadPassword(0)
		if err != nil {
		}
		return Account{Username: username, Password: string(password)}, nil
	}
	//if config file exists
	viper.SetConfigName("instago")
	viper.SetConfigType("env")
	viper.AddConfigPath("$HOME/.instago.env")
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
		} else {
			// Config file was found but another error was produced
		}
	}
	username := viper.GetString("USERNAME")
	password := viper.GetString("PASSWORD")

	return Account{Username: username, Password: password}, nil
}

func login(a *Account) {

}