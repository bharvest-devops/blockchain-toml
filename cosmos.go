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
		return err
	}
	return nil
}

func (c *CosmosConfigFile) ExportMergeWithDefault() ([]byte, error) {
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

func (c *CosmosConfigFile) MergeWithConfig(o CosmosConfigFile) error {
	if err := mergo.Merge(c, o, mergo.WithoutDereference, mergo.WithOverwriteWithEmptyValue); err != nil {
		return err
	}
	return nil
}

func (c *CosmosConfigFile) ExportMergeWithConfig(o CosmosConfigFile) ([]byte, error) {
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
		return err
	}
	return nil
}

func (c *CosmosAppFile) ExportMergeWithDefault() ([]byte, error) {
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

func (c *CosmosAppFile) MergeWithConfig(o *CosmosAppFile) error {
	if err := mergo.Merge(c, o, mergo.WithoutDereference); err != nil {
		return err
	}
	return nil
}

func (c *CosmosAppFile) ExportMergeWithConfig(o *CosmosAppFile) ([]byte, error) {
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
