package blockchain_toml

import (
	"dario.cat/mergo"
	_ "embed"
	"github.com/pelletier/go-toml"
	"log"
)

var (
	//go:embed toml/config.toml
	defaultNamadaConfigTomlByte []byte
)

func getDefaultNamadaConfig() *Config {
	var defaultConfig Config
	err := toml.Unmarshal(defaultNamadaConfigTomlByte, &defaultConfig)
	if err != nil {
		panic("Cannot convert config.toml into structure!!" + err.Error())
	}
	return &defaultConfig
}

func ExportMergedNamadaConfigBytes(config *Config) ([]byte, error) {
	if err := MergeNamadaConfigToml(config); err != nil {
		log.Fatalf("Merge failed! \n%s", err.Error())
	}
	return toml.Marshal(config)
}

func MergeNamadaConfigToml(config *Config) error {
	if err := mergo.Merge(config, *getDefaultNamadaConfig(), mergo.WithoutDereference); err != nil {
		return err
	}
	return nil
}
