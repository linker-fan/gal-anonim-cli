package config

import (
	"log"
	"os"

	"github.com/linker-fan/gal-anonim-cli/internal/utils"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerAPI struct {
		Host             string `yaml:"host"`
		Port             string `yaml:"port"`
		RegisterEndpoint string `yaml:"registerEndpoint"`
		LoginEndpoint    string `yaml:"loginEndpoint"`
	} `yaml:"serverAPI"`
	Local struct {
		TokeFilePath string `yaml:"tokenFilePath"`
	} `yaml:"local"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}

	err := utils.ValidatePath(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
