package config

type StoreConfig struct {
	Type string `json:"type" yaml:"type"`
}

func newStoreConfig() *StoreConfig {
	return &StoreConfig{
		Type: "memory",
	}
}
