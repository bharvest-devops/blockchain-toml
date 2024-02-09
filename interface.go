package blockchain_toml

type CometBFTFile interface {
	MergeWithDefault() error
	ExportMergeWithDefault() ([]byte, error)
	MergeWithConfig(o *CometBFTFile) error
	ExportMergeWithConfig(o *CometBFTFile) ([]byte, error)
}
