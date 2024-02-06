package blockchain_toml

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("addConfigToml test", AddConfigTomlTest)
	t.Run("ExportMergedNamadaConfigFile test", ExportMergedConfigTest)
}

func AddConfigTomlTest(t *testing.T) {
	var config NamadaConfigFile
	WASM_DIR := "Hello"
	config.WasmDir = &WASM_DIR

	err := MergeNamadaConfigFileToml(&config)
	if err != nil {
		t.Fatalf("Error occurred: %s", err.Error())
		return
	} else {
		t.Log("Success")
		return
	}
}

func ExportMergedConfigTest(t *testing.T) {
	var config NamadaConfigFile
	WASM_DIR := "ExportMergedNamadaConfigFile"
	config.WasmDir = &WASM_DIR

	bytes, err := ExportMergedNamadaConfigFileBytes(&config)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}
