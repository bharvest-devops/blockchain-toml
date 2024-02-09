package blockchain_toml

import (
	"testing"
)

func TestNamada(t *testing.T) {
	t.Run("namada merge test", testNamadaConfigMerge)
	t.Run("namada export merge test", testNamadaConfigExportMerge)
}

func testNamadaConfigMerge(t *testing.T) {
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

func testNamadaConfigExportMerge(t *testing.T) {
	var config NamadaConfigFile
	WASM_DIR := "ExportMergedNamadaConfigFile"
	config.WasmDir = &WASM_DIR

	bytes, err := config.ExportMergeWithDefault()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}
