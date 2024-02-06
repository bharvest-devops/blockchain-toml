package blockchain_toml

import (
	"testing"
)

func CosmosConfigMergeTest(t *testing.T) {
	var config CosmosConfigFile
	config.BoolBlockSync = false

	err := config.Merge()
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
	config.BoolBlockSync = false

	bytes, err := config.ExportMerge()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}

func CosmosAppMergeTest(t *testing.T) {
	var config CosmosAppFile
	config.PruningInterval = "0"

	err := config.Merge()
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
	config.PruningInterval = "0"

	bytes, err := config.ExportMerge()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}
