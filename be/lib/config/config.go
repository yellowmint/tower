package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var configuration Config

func Init() {
	viper.SetEnvPrefix("TOWER")
	viper.AutomaticEnv()

	err := viper.BindEnv("port", "PORT")
	if err != nil {
		panic(fmt.Errorf("cannot bind 'PORT' environment varirable: %w\n", err))
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("lib/config/")
	viper.SetConfigName("default")
	loadConfig()

	switch getMode() {
	case Prod:
		viper.SetConfigName("prod")
	case Dev:
		viper.SetConfigName("dev")
	case IntegrationTest:
		viper.SetConfigName("integration-test")
	}

	loadConfig()
}

func getMode() mode {
	switch viper.GetString("mode") {
	case "prod":
		return Prod
	case "integrationTest":
		return IntegrationTest
	default:
		return Dev
	}
}

func loadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot read config file: %w\n", err))
	}

	err = viper.UnmarshalExact(&configuration)
	if err != nil {
		panic(fmt.Errorf("cannot unmarshal configuration: %w\n", err))
	}
}

func Get() Config {
	if configuration.Version == "" {
		panic("configuration not loaded")
	}

	return configuration
}
