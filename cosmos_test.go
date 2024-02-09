package blockchain_toml

import (
	"testing"
)

func CosmosConfigMergeTest(t *testing.T) {
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

func CosmosConfigExportMergeTest(t *testing.T) {
	var config CosmosConfigFile
	boolBlockSync := false
	config.BoolBlockSync = &boolBlockSync

	bytes, err := config.ExportMergeWithDefault()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}

func CosmosAppMergeTest(t *testing.T) {
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

func CosmosAppExportMergeTest(t *testing.T) {
	var config CosmosAppFile
	zero := "0"
	config.PruningInterval = &zero

	bytes, err := config.ExportMergeWithDefault()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}
