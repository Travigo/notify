package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var Global *Config

type Config struct {
	Version string

	Providers []*ProviderConfig
	Events    []*EventConfig
}

func (c *Config) Verify() error {
	if c.Version != "1" {
		return errors.New(fmt.Sprintf("Version must be 1, but got %s", c.Version))
	}

	return nil
}

func LoadFromFile(path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config file")
	}

	config := Config{}

	err = yaml.Unmarshal([]byte(file), &config)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshal config file")
	}

	err = config.Verify()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to verify config file")
	}

	Global = &config
}
