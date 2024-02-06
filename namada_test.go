package blockchain_toml

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("addConfigToml test", AddConfigTomlTest)
	t.Run("ExportMergedConfig test", ExportMergedConfigTest)
}

func AddConfigTomlTest(t *testing.T) {
	var config Config
	WASM_DIR := "Hello"
	config.WasmDir = &WASM_DIR

	err := AddConfigToml(&config)
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
	WASM_DIR := "ExportMergedConfig"
	config.WasmDir = &WASM_DIR

	bytes, err := ExportMergedConfig(&config)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	t.Logf("\n%s", string(bytes))

}
