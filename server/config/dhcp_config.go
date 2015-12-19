package config

type DhcpConfig struct {
	Type        string `json:"type" yaml:"type"`
	DnsmasqFile string `json:"dnsmasq_file" yaml:"dnsmasq_file"`
}

func newDhcpConfig() *DhcpConfig {
	return &DhcpConfig{
		Type:        "dnsmasq",
		DnsmasqFile: "/var/lib/misc/dnsmasq.leases",
	}
}
