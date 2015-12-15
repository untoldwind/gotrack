package config

type ProviderConfig struct {
	Type     string `json:"type" yaml:"type"`
	ProcFile string `json:"proc_file" yaml:"proc_file"`
}

func newProvderConfig() *ProviderConfig {
	return &ProviderConfig{
		Type:     "proc",
		ProcFile: "/proc/net/ip_conntrack",
	}
}

func readProviderConfig(fileName string) (*ProviderConfig, error) {
	var providerConfig ProviderConfig

	if err := loadConfigFile(fileName, &providerConfig); err != nil {
		return nil, err
	}

	return &providerConfig, nil
}
