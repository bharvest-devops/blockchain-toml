package blockchain_toml

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("addConfigToml test", AddConfigTomlTest)
	t.Run("ExportMergedNamadaConfig test", ExportMergedConfigTest)
}

func AddConfigTomlTest(t *testing.T) {
	var config Config
	WASM_DIR := "Hello"
	config.WasmDir = &WASM_DIR

	err := MergeNamadaConfigToml(&config)
	if err != nil {
		t.Fatalf("Error occurred: %s", err.Error())
		return
	} else {
		t.Log("Success")
		return
	}
}

func ExportMergedConfigTest(t *testing.T) {
	var config Config
	WASM_DIR := "ExportMergedNamadaConfig"
	config.WasmDir = &WASM_DIR

	bytes, err := ExportMergedNamadaConfigBytes(&config)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}
