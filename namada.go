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

func ExportMergedConfig(config *Config) ([]byte, error) {
	if err := AddConfigToml(config); err != nil {
		log.Fatalf("Merge failed! \n%s", err.Error())
	}
	return toml.Marshal(config)
}

func AddConfigToml(config *Config) error {
	if err := mergo.Merge(config, *getDefaultConfig(), mergo.WithoutDereference); err != nil {
		return err
	}
	return nil
}
