package config

type ProviderConfig struct {
	Type          string `json:"type" yaml:"type"`
	ConntrackFile string `json:"proc_file" yaml:"proc_file"`
	DevFile       string `json:"dev_file" yaml:"dev_file"`
	WanInterface  string `json:"wan_interface" yaml:"wan_interface"`
}

func newProviderConfig() *ProviderConfig {
	return &ProviderConfig{
		Type:          "proc",
		ConntrackFile: "/proc/net/ip_conntrack",
		DevFile:       "/proc/net/dev",
		WanInterface:  "eth0",
	}
}

func readProviderConfig(fileName string) (*ProviderConfig, error) {
	var providerConfig ProviderConfig

	if err := loadConfigFile(fileName, &providerConfig); err != nil {
		return nil, err
	}

	return &providerConfig, nil
}
