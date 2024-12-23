package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `config:"server"`
	Database DatabaseConfig `config:"database"`
}

type ServerConfig struct {
	Port int    `config:"port"`
	URL  string `config:"url"`
}

type DatabaseConfig struct {
	User         string `config:"user"`
	Pass         string `config:"pass"`
	Port         int    `config:"port"`
	MaxIdleConn  int    `config:"max_idle_conn"`
	MaxOpenConn  int    `config:"max_open_conn"`
	DefaultLimit int    `config:"default_limit"`
}

var appConfig Config

func Init() {
	// Set the file name of the configuration file
	viper.SetConfigName("config") // Name without extension
	viper.SetConfigType("yaml")   // Explicitly set the file type
	viper.AddConfigPath(".")      // Look for the file in the current directory

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	// Map Viper config to the struct
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatalf("Error unmarshaling config, %v", err)
	}

	fmt.Printf("App Name: %s\n", appConfig.Server.Port)
}

func GetAppConfig() Config {
	return appConfig
}