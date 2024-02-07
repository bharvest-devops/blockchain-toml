package blockchain_toml

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("namada merge test", NamadaConfigMergeTest)
	t.Run("namada export merge test", NamadaConfigExportMergeTest)

	t.Run("cosmos merge test", CosmosConfigMergeTest)
	t.Run("cosmos export merge test", CosmosConfigExportMergeTest)

	t.Run("cosmos merge test", CosmosAppMergeTest)
	t.Run("cosmos export merge test", CosmosAppExportMergeTest)
}

func NamadaConfigMergeTest(t *testing.T) {
	var config NamadaConfigFile
	WASM_DIR := "Hello"
	config.WasmDir = &WASM_DIR

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

func NamadaConfigExportMergeTest(t *testing.T) {
	var config NamadaConfigFile
	WASM_DIR := "ExportMergedNamadaConfigFile"
	config.WasmDir = &WASM_DIR

	bytes, err := config.ExportMergeWithDefault()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}
