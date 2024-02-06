package namada_toml

import (
	_ "embed"
	"github.com/pelletier/go-toml"
)

type decodedToml = map[string]any

var (
	//go:embed toml/config.toml
	defaultConfigTomlByte []byte
)

func init() {
	var defaultConfig Config
	err := toml.Unmarshal(defaultConfigTomlByte, &defaultConfig)
	if err != nil {
		panic("Cannot convert config.toml into structure!!" + err.Error())
	}
}

func BuildConfigToml(overrides *string) error {
	return nil
}

func getDefaultConfig() *Config {

	return nil
}

func addConfigToml(config *Config) error {
	//if config.WasmDir == nil {
	//	config.WasmDir = defaultConfig.WasmDir
	//}

	return nil
}
