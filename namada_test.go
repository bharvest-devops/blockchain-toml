package namada_toml

import "testing"

func Test(t *testing.T) {
	t.Run("addConfigToml test", AddConfigTomlTest)
}

func AddConfigTomlTest(t *testing.T) {
	var config Config

	_ = addConfigToml(&config)
	//if err != nil {
	//	t.Fatal("error occurred")
	//	return
	//} else {
	//	t.Logf("Error didn't occur\n content: %s", *config.WasmDir)
	//}
}
