package config

type StoreConfig struct {
	Type    string `json:"type" yaml:"type"`
	LanCIDR string `json:"lan_cidr" yaml:"lan_cidr"`
}

func newStoreConfig() *StoreConfig {
	return &StoreConfig{
		Type:    "memory",
		LanCIDR: "192.168.1.0/24",
	}
}
