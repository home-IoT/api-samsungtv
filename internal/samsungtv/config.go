package samsungtv

import (
	"io/ioutil"
	"os"

	"github.com/home-IoT/api-samsungtv/gen/restapi/operations"
	"github.com/home-IoT/api-samsungtv/internal/log"
	"gopkg.in/yaml.v2"
)

const defaultControllerName = "SamsungTVController"

const protocolWS = "ws"
const protocolWSS = "wss"

var defaultTVProtocol = protocolWSS

var defaultPortWS = "8001"
var defaultPortWSS = "8002"
var defaultTVPort = defaultPortWS

type tvConfig struct {
	Host     string  `yaml:"host"`
	Mac      *string `yaml:"mac"`
	Port     *string `yaml:"port"`
	Protocol *string `yaml:"protocol"`
}

type controllerConfig struct {
	Name string `yaml:"name"`
}

type samsungTVConfigYAML struct {
	TV         tvConfig         `yaml:"tv"`
	Controller controllerConfig `yaml:"controller"`
}

var configuration *samsungTVConfigYAML

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

	config := new(samsungTVConfigYAML)

	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Debugf("%v", err)
		log.Exitf(1, "Error loading the configuration file.")
	}

	setDefaultValues(config)

	log.InitLog(true)

	configuration = config
}

func setDefaultValues(config *samsungTVConfigYAML) {
	if config.TV.Protocol == nil {
		config.TV.Protocol = &defaultTVProtocol
	}

	if config.TV.Port == nil {
		switch *config.TV.Protocol {
		case protocolWS:
			config.TV.Port = &defaultPortWS
		case protocolWSS:
			config.TV.Port = &defaultPortWSS
		default:
			config.TV.Port = &defaultTVPort
		}
	}

	if len(config.Controller.Name) == 0 {
		config.Controller.Name = defaultControllerName
	}
}
