package app

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type (
	AppConfig struct {
		ApiKey    string `yaml:"api_key"`
		ApiSecret string `yaml:"api_secret"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
	}
	ConfigFile map[string]*AppConfig
)

func LoadConfig(env string) (*AppConfig, error) {
	configFile := ConfigFile{}
	file, _ := os.Open("config.yaml")
	defer file.Close()
	decoder := yaml.NewDecoder(file)

	// Always check for errors!
	if err := decoder.Decode(&configFile); err != nil {
		println(err.Error())
		return nil, err
	}

	println("before env")
	appConfig, ok := configFile[env]
	if !ok {
		return nil, fmt.Errorf("no such environment: %s", env)
	}

	return appConfig, nil
}
