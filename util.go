package blockchain_toml

import "fmt"

func printError(err error) error {
	return fmt.Errorf("toml task failed %s", err.Error())
}
