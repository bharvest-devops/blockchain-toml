package namada_toml

import (
	"dario.cat/mergo"
	_ "embed"
	"github.com/pelletier/go-toml"
	"log"
)

type decodedToml = map[string]any

var (
	//go:embed toml/config.toml
	defaultConfigTomlByte []byte
)

func getDefaultConfig() *Config {
	var defaultConfig Config
	err := toml.Unmarshal(defaultConfigTomlByte, &defaultConfig)
	if err != nil {
		panic("Cannot convert config.toml into structure!!" + err.Error())
	}
	return &defaultConfig
}

func AddConfigToml(config *Config) error {
	if err := mergo.Merge(config, getDefaultConfig(), mergo.WithoutDereference); err != nil {
		log.Fatalf("Cannot merge with default config!\n %s", err.Error())
	}

	return nil
}
