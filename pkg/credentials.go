package pkg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

//TODO reformat
var username string
var password string
var answer string

//Account data struct stores account credentials and metrics
type Account struct {
	//TODO: add instagram golang api
	Username  string
	Password  string
	Auth      bool // successful login status
	Following map[string]bool
	Followers map[string]bool
	Ghosts    []string //users who dont follow back
}

//getUserPass gets username and password via user input
func getUserPass() (string, string, error) {
	fmt.Scanln("Username: ")
	fmt.Scanln(&username)
	password, err := terminal.ReadPassword(0)
	if err != nil {
		//TODO
	}
	viper.Set("USERNAME", username)
	viper.Set("PASSWORD", password)
	return username, string(password), nil
}

// createConfig creates ~/.instago.env file when called
func createConfig() (Account, error) {
	username, password, err := getUserPass()
	if err != nil {
		//TODO
	}
	viper.SetDefault("USERNAME", "")
	viper.SetDefault("PASSWORD", "")
	fmt.Println("account credentials stored in ~/.instago.env")
	return Account{Username: username, Password: string(password)}, nil
}

//yesNoUsrInputForConfig gets credentials from user input and creates config file depending on y/n answer
func yesNoUsrInputForConfig(answer string) (Account, error) {
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		credentials, err := createConfig()
		return credentials, err
	}
	// n- get user & pass via user input
	username, password, err := getUserPass()
	if err != nil {
		//TODO
	}
	//TODO if answer != y/n, print srry. don't understand
	return Account{Username: username, Password: string(password)}, nil
}

//GetUserCredentials gets the IG username and password from a config file or directly from a user's input
func GetUserCredentials() (Account, error) {
	viper.SetConfigName("instago")
	viper.SetConfigType("env")
	viper.AddConfigPath("$HOME/.instago.env")
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		//when config not found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Scanln("~/.instago.env file does not exist. Would you like to create one? (Y/N) ")
			fmt.Scanln(&answer)
			credentials, err := yesNoUsrInputForConfig(answer)
			return credentials, err
		} else {
			// Config file was found but another error was produced
			fmt.Scanln("unable to parse config file. would you like to overwrite with a new config? Y/N ")
			fmt.Scanln(&answer)
			credentials, err := yesNoUsrInputForConfig(answer)
			return credentials, err
		}
	}
	//if config file exists and read with no errors
	username := viper.GetString("USERNAME")
	password := viper.GetString("PASSWORD")
	//TODO when wrong username/password combo used -- when cannot login with credentials from config
	return Account{Username: username, Password: password}, nil
}

func login(a *Account) {

}
