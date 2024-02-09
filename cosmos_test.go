package blockchain_toml

import (
	"testing"
)

func TestCosmos(t *testing.T) {
	t.Run("cosmos config merge test", testCosmosConfigMerge)
	t.Run("cosmos config export merge test", testCosmosConfigExportMerge)

	t.Run("cosmos app merge test", testCosmosAppMerge)
	t.Run("cosmos app export merge test", testCosmosAppExportMerge)
	t.Run("cosmos app merge with tomlOverrides", testCosmosAppMergeWithTomlOverrides)
}
func testCosmosConfigMerge(t *testing.T) {
	var config CosmosConfigFile
	boolBlockSync := false
	config.BoolBlockSync = &boolBlockSync

	err := config.MergeWithDefault()
	if err != nil {
		return
	}
	if err != nil {
		t.Fatalf("Error occurred: %s", err.Error())
		return
	} else {
		t.Log("Success")
		return
	}
}

func testCosmosConfigExportMerge(t *testing.T) {
	var config CosmosConfigFile
	boolBlockSync := false
	config.BoolBlockSync = &boolBlockSync

	bytes, err := config.ExportMergeWithDefault()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}

func testCosmosAppMerge(t *testing.T) {
	var config CosmosAppFile
	zero := "0"
	config.PruningInterval = &zero

	err := config.MergeWithDefault()
	if err != nil {
		return
	}
	if err != nil {
		t.Fatalf("Error occurred: %s", err.Error())
		return
	} else {
		t.Log("Success")
		return
	}
}

func testCosmosAppExportMerge(t *testing.T) {
	var config CosmosAppFile
	zero := "0"
	config.PruningInterval = &zero

	bytes, err := config.ExportMergeWithDefault()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}

func testCosmosAppMergeWithTomlOverrides(t *testing.T) {
	var config CosmosAppFile
	overrides := `
	minimum-gas-prices = "0.1override"
	new-base = "new base value"
	
	[api]
	enable = false
	new-field = "test"
	`
	bytes, err := config.ExportMergeWithTomlOverrides([]byte(overrides))
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))
}
