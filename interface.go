package blockchain_toml

type TomlFile interface {
	MergeWithDefault() error
	ExportMergeWithDefault() ([]byte, error)
	MergeWithConfig(o *TomlFile) error
	ExportMergeWithConfig(o *TomlFile) ([]byte, error)
}
