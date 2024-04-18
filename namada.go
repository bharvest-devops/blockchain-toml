package blockchain_toml

import (
	"dario.cat/mergo"
	_ "embed"
	"github.com/pelletier/go-toml/v2"
)

var (
	//go:embed toml/namadaConfig.toml
	defaultNamadaConfigFileTomlByte []byte
)

func getDefaultNamadaConfigFile() *NamadaConfigFile {
	var defaultFile NamadaConfigFile
	err := toml.Unmarshal(defaultNamadaConfigFileTomlByte, &defaultFile)
	if err != nil {
		panic("Cannot convert NamadaConfigFile.toml into structure!!" + err.Error())
	}
	return &defaultFile
}

func (c *NamadaConfigFile) MergeWithDefault() error {
	if err := mergo.Merge(c, *getDefaultNamadaConfigFile()); err != nil {
		return printError(err)
	}
	return nil
}

func (c *NamadaConfigFile) ExportMergeWithDefault() ([]byte, error) {
	err := c.MergeWithDefault()
	if err != nil {
		return nil, printError(err)
	}
	t, err := toml.Marshal(c)
	if err != nil {
		return nil, printError(err)
	}
	return t, nil
}

func (c *NamadaConfigFile) MergeWithConfig(o *NamadaConfigFile) error {
	if err := mergo.Merge(c, o); err != nil {
		return printError(err)
	}
	return nil
}

func (c *NamadaConfigFile) ExportMergeWithConfig(o *NamadaConfigFile) ([]byte, error) {
	err := c.MergeWithConfig(o)
	if err != nil {
		return nil, printError(err)
	}
	t, err := toml.Marshal(c)
	if err != nil {
		return nil, printError(err)
	}
	return t, nil
}

func (c *NamadaConfigFile) ExportMergeWithTomlOverrides(overrides []byte) ([]byte, error) {
	var (
		overridesMap map[string]interface{}
		originMap    map[string]interface{}
		err          error
	)
	if err = toml.Unmarshal(overrides, &overridesMap); err != nil {
		return nil, printError(err)
	}
	tBytes, err := toml.Marshal(c)
	err = toml.Unmarshal(tBytes, &originMap)
	if err != nil {
		return nil, printError(err)
	}
	if err = mergo.Merge(&originMap, overridesMap); err != nil {
		return nil, printError(err)
	}
	resp, err := toml.Marshal(originMap)
	if err != nil {
		return nil, printError(err)
	}
	return resp, nil
}
