package blockchain_toml

import (
	"dario.cat/mergo"
	_ "embed"
	"github.com/pelletier/go-toml"
)

var (
	//go:embed toml/cosmosConfig.toml
	defaultCosmosConfigFileTomlByte []byte

	//go:embed toml/cosmosApp.toml
	defaultCosmosAppFileTomlByte []byte
)

func getDefaultCosmosConfigFile() *CosmosConfigFile {
	var defaultFile CosmosConfigFile
	err := toml.Unmarshal(defaultCosmosConfigFileTomlByte, &defaultFile)
	if err != nil {
		panic("Cannot convert NamadaConfigFile.toml *into structure!!" + err.Error())
	}
	return &defaultFile
}

func (c *CosmosConfigFile) MergeWithDefault() error {
	if err := mergo.Merge(c, *getDefaultCosmosConfigFile(), mergo.WithoutDereference); err != nil {
		return printError(err)
	}
	return nil
}

func (c *CosmosConfigFile) ExportMergeWithDefault() ([]byte, error) {
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

func (c *CosmosConfigFile) MergeWithConfig(o CosmosConfigFile) error {

	if err := mergo.Merge(c, o, mergo.WithoutDereference, mergo.WithOverwriteWithEmptyValue); err != nil {
		return printError(err)
	}
	return nil
}

func (c *CosmosConfigFile) ExportMergeWithConfig(o CosmosConfigFile) ([]byte, error) {
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

func (c *CosmosConfigFile) ExportMergeWithTomlOverrides(overrides []byte) ([]byte, error) {
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
	if err = mergo.Merge(&originMap, overridesMap, mergo.WithoutDereference, mergo.WithOverride); err != nil {
		return nil, printError(err)
	}
	resp, err := toml.Marshal(originMap)
	return resp, printError(err)
}

func getDefaultCosmosAppFile() *CosmosAppFile {
	var defaultFile CosmosAppFile
	err := toml.Unmarshal(defaultCosmosAppFileTomlByte, &defaultFile)
	if err != nil {
		panic("Cannot convert NamadaConfigFile.toml *into structure!!" + err.Error())
	}
	return &defaultFile
}

func (c *CosmosAppFile) MergeWithDefault() error {
	if err := mergo.Merge(c, *getDefaultCosmosAppFile(), mergo.WithoutDereference); err != nil {
		return printError(err)
	}
	return nil
}

func (c *CosmosAppFile) ExportMergeWithDefault() ([]byte, error) {
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

func (c *CosmosAppFile) MergeWithConfig(o *CosmosAppFile) error {
	if err := mergo.Merge(c, o, mergo.WithoutDereference); err != nil {
		return printError(err)
	}
	return nil
}

func (c *CosmosAppFile) ExportMergeWithConfig(o *CosmosAppFile) ([]byte, error) {
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

func (c *CosmosAppFile) ExportMergeWithTomlOverrides(overrides []byte) ([]byte, error) {
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
	if err = mergo.Merge(&originMap, overridesMap, mergo.WithoutDereference, mergo.WithOverride); err != nil {
		return nil, printError(err)
	}
	return toml.Marshal(originMap)
}
