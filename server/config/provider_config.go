package config

type ProviderConfig struct {
	Type          string `json:"type" yaml:"type"`
	ConntrackFile string `json:"proc_file" yaml:"proc_file"`
	DevFile       string `json:"dev_file" yaml:"dev_file"`
	WanInterface  string `json:"wan_interface" yaml:"wan_interface"`
	LanInterface  string `json:"lan_interface" yaml:"lan_interface"`
}

func newProviderConfig() *ProviderConfig {
	return &ProviderConfig{
		Type:          "proc",
		ConntrackFile: "/proc/net/ip_conntrack",
		DevFile:       "/proc/net/dev",
		WanInterface:  "eth0",
		LanInterface:  "eth1",
	}
}
