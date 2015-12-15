package config

type StoreConfig struct {
	Type string `json:"type" yaml:"type"`
}

func newStoreConfig() *StoreConfig {
	return &StoreConfig{}
}

func readStoreConfig(fileName string) (*StoreConfig, error) {
	var storeConfig StoreConfig

	if err := loadConfigFile(fileName, &storeConfig); err != nil {
		return nil, err
	}

	return &storeConfig, nil
}
