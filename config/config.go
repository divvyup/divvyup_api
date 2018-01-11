package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

/*
	Config is our type to encapsulate everything the other
	parts of our app will need
*/
type config struct {
	Secret    string `json:"secret"`
	APIPrefix string `json:"api_prefix"`
}

var globalConfig config

/*
AppConfig reads our environment variables and configures the following:
	1. production vs development mode
	2. database stuff
*/
func AppConfig() {
	env := os.Getenv("DIVVYUP_API_MODE")

	if env == "" {
		fmt.Println("No divvyup environment variable specified.")
		os.Exit(2)
	}

	if env == "development" {
		fmt.Printf("Starting divvyup api in %s mode...\n", color.GreenString("DEVELOPMENT"))

		// Read and then parse our config file
		conf, err := ioutil.ReadFile("config/dev.json")

		if err != nil {
			fmt.Println("Error reading development config file.")
			os.Exit(2)
		}

		// Parse out some useful information
		jsone := json.Unmarshal(conf, &globalConfig)

		if jsone != nil {
			fmt.Println("Configuration file does not have the proper structure.")
			os.Exit(2)
		}

		fmt.Printf("Configured to use %s as the API prefix\n", color.RedString(globalConfig.APIPrefix))

	} else if env == "production" {
		fmt.Printf("Starting divvyup api in %s mode...\n", color.GreenString("PRODUCTION"))
	} else {
		fmt.Println("Invalid api mode given.")
		os.Exit(2)
	}

}

/*
	Secret will allow other parts of the application to acccess our
	secret. This will be used for signing jwt's etc.
*/
func Secret() string {
	return globalConfig.Secret
}

/*
	Prefix will allow our routing to be under the proper api versioning
	scheme
*/
func Prefix() string {
	return globalConfig.APIPrefix
}
