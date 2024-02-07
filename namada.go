package blockchain_toml

import (
	"dario.cat/mergo"
	_ "embed"
	"github.com/pelletier/go-toml"
)

var (
	//go:embed toml/namadaConfig.toml
	defaultNamadaConfigFileTomlByte []byte
)

func getDefaultNamadaConfigFile() *NamadaConfigFile {
	var defaultFile NamadaConfigFile
	err := toml.Unmarshal(defaultNamadaConfigFileTomlByte, &defaultFile)
	if err != nil {
		panic("Cannot convert NamadaConfigFile.toml *into structure!!" + err.Error())
	}
	return &defaultFile
}

func (c *NamadaConfigFile) ExportMerge() ([]byte, error) {
	err := c.Merge()
	if err != nil {
		return nil, err
	}
	t, err := toml.Marshal(c)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *NamadaConfigFile) Merge() error {
	if err := mergo.Merge(c, *getDefaultNamadaConfigFile(), mergo.WithoutDereference); err != nil {
		return err
	}
	return nil
}
