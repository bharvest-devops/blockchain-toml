package blockchain_toml

type TomlFile *interface {
	MergeWithDefault() error
	ExportMergeWithDefault() ([]byte, error)
	MergeWithConfig(c *TomlFile) error
	ExportMergeWithConfig(c *TomlFile) ([]byte, error)
}
