package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigDatabase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func Config(path string) (config ConfigDatabase, err error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return config, err
	}
	defer f.Close()
	return config, yaml.NewDecoder(f).Decode(&config)
}
