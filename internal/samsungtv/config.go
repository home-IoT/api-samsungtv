package samsungtv

import (
	"io/ioutil"
	"os"

	"github.com/home-IoT/api-samsungtv/gen/restapi/operations"
	"github.com/home-IoT/api-samsungtv/internal/log"
	"gopkg.in/yaml.v2"
)

type receiverConfig struct {
	Host string `yaml:"host"`
}

type denonConfigYAML struct {
	TV receiverConfig `yaml:"tv"`
}

var configuration *denonConfigYAML

// Configure configures the server with a given configuration file
func Configure(api *operations.SamsungtvAPI) {
	options := getConfigurationOptions(api)

	if options.Version {
		showVersion()
		os.Exit(0)
	}

	if options.ConfigFile == "" {
		printError("Configuration file is missing. Use flag `-c, --config' to provide a config file.")
		os.Exit(1)
	}
	loadConfig(options.ConfigFile)
}

func loadConfig(configFile string) {

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Debugf("%v", err)
		log.Exitf(1, "Error loading the configuration file.")
	}

	config := new(denonConfigYAML)

	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Debugf("%v", err)
		log.Exitf(1, "Error loading the configuration file.")
	}

	configuration = config
}
