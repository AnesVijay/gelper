package config

import (
	"fmt"
	"os"

	"github.com/AnesVijay/glogger"
	"gopkg.in/yaml.v3"
)

var Conf *Config

func New(pathToFile string) *Config {
	var conf Config

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		glogger.GetLogger().SendError(fmt.Sprintf("failed to read config file (%s): %v", pathToFile, err))
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		glogger.GetLogger().SendError(fmt.Sprintf("failed unmarshaling YAML: %v", err))
	}

	Conf = &conf

	return Conf
}
