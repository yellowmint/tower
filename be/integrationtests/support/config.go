package support

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"regexp"
)

func loadConfig() {
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(getIntegrationTestRootDir() + "support/")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot read config file: %w\n", err))
	}
}

func getIntegrationTestRootDir() string {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("cannot get current working directory: %w\n", err))
	}

	rootChildren := regexp.MustCompile(`^(.+/be/integrationtests)(.+)$`)

	return rootChildren.ReplaceAllString(workingDir, "$1") + "/"
}
