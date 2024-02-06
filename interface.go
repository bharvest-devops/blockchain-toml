package blockchain_toml

type TomlFile interface {
	Merge() error
	ExportMerge() ([]byte, error)
}
