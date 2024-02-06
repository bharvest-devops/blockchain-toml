package namada_toml

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("addConfigToml test", AddConfigTomlTest)
}

func AddConfigTomlTest(t *testing.T) {
	var config Config
	hello := "Hello"
	config.WasmDir = &hello

	err := AddConfigToml(&config)
	if err != nil {
		t.Fatalf("Error occurred: %s", err.Error())
		return
	} else {
		t.Log("Success")
		return
	}
}
