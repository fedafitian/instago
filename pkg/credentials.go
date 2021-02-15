package pkg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	username string
	password string
	answer string
)

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
	//TODO: decide to ask for "username" or "handle"
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		//TODO
	}
	return username, string(password), nil
}

// createConfig creates ~/.instago.env file when called
func createConfig(cfgFile, username, password string) (Account, error) {
	//do i need setdefault ?
	viper.SetConfigFile(cfgFile)
	viper.Set("USERNAME", username)
	viper.Set("PASSWORD", password)
	viper.WriteConfigAs(cfgFile)
	fmt.Printf("account credentials stored in %s", cfgFile)
	return Account{Username: username, Password: string(password)}, nil
}

//yesNoUsrInputForConfig gets credentials from user input and creates config file depending on y/n answer
func yesNoUsrInputForConfig(cfgFile, answer string) (Account, error) {
	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		username, password, err := getUserPass()
		if err != nil {

		}
		credentials, err := createConfig(cfgFile, username, password)
		return credentials, err
	}
	// n- get user & pass via user input
	username, password, err := getUserPass()
	if err != nil {
		//TODO
	}
	//TODO if answer != y/n, print srry. don't understand
	return Account{Username: username, Password: password}, nil
}

//GetUserCredentials gets the IG username and password from a config file or directly from a user's input
func GetUserCredentials(cfgFile string) (Account, error) {
	//TODO: figure out why default config file path not recognized
	// viper.SetConfigName("instago")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("$HOME/.instago.")
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		//when config not found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Print("~/.instago.env file does not exist. Would you like to create one? (Y/N):  ")
			fmt.Scanln(&answer)
			credentials, err := yesNoUsrInputForConfig(cfgFile, answer)
			return credentials, err
		} else {
			// Config file was found but another error was produced
			fmt.Print("unable to parse config file. would you like to overwrite with a new config? (Y/N): ")
			fmt.Scanln(&answer)
			credentials, err := yesNoUsrInputForConfig(cfgFile, answer)
			return credentials, err
		}
	}
	//if config file exists and read with no errors
	username = viper.GetString("USERNAME")
	password = viper.GetString("PASSWORD")
	// condition when username and password not set
	if username == "" && password == ""{
		fmt.Println("config found but missing username & password. provide details to add ")
		username, password, err := getUserPass()
		if err != nil {
			//TODO
		}
		createConfig(cfgFile, username, password)
	// condition when username not set
	} else if username == "" && password != "" {
		fmt.Print("username not found in config. what is the username?: ")
		fmt.Scanln(&username)
		//TODO: decide to update password here
	// condition when password not set
	} else if username != "" && password == ""{
		fmt.Printf("password not found in config. what is the password for %s?: ", username)
		password, err := terminal.ReadPassword(0)
		if err != nil {
			//TODO
		}
		viper.Set("PASSWORD", string(password))
		//TODO: keep default order in env file
		createConfig(cfgFile, username, string(password))
	}

	fmt.Printf("%s config file found for %s account", cfgFile, username)

	//TODO when wrong username/password combo used -- when cannot login with credentials from config --> see viper.SetEnvKeyReplacer
	return Account{Username: username, Password: password}, nil
}

func login(a *Account) {

}
