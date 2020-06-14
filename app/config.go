package app

import (
	"encoding/json"
	"os"
)

type Config struct {
	API     API    `json:"api"`
	AppName string `json:"appName"`
}

type API struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewConfigFile(filename string) error {
	err := generateConfigFile(filename, configSample())
	if err != nil {
		return err
	}
	return nil
}

func generateConfigFile(filename string, config Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func configSample() Config {
	var c Config
	return c
}
