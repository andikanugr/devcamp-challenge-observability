package config

import (
	"os"

	"github.com/spf13/viper"
)

func Init() (*Schema, error) {
	env := os.Getenv("ENV")
	var configPath string
	if env == "production" {
		configPath = "/app/files/config/config.prod.yml"
	} else {
		// Default to development environment
		configPath = "./files/config/config.dev.yml"
	}

	// Set the file name of the configurations file
	viper.SetConfigFile(configPath)

	// Read and unmarshal the config file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Schema
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
