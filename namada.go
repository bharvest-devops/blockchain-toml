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

func (c *NamadaConfigFile) MergeWithDefault() error {
	if err := mergo.Merge(c, *getDefaultNamadaConfigFile(), mergo.WithoutDereference); err != nil {
		return err
	}
	return nil
}

func (c *NamadaConfigFile) ExportMergeWithDefault() ([]byte, error) {
	err := c.MergeWithDefault()
	if err != nil {
		return nil, err
	}
	t, err := toml.Marshal(c)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *NamadaConfigFile) MergeWithConfig(o *NamadaConfigFile) error {
	if err := mergo.Merge(c, o, mergo.WithoutDereference); err != nil {
		return err
	}
	return nil
}

func (c *NamadaConfigFile) ExportMergeWithConfig(o *NamadaConfigFile) ([]byte, error) {
	err := c.MergeWithConfig(o)
	if err != nil {
		return nil, err
	}
	t, err := toml.Marshal(c)
	if err != nil {
		return nil, err
	}
	return t, nil
}
