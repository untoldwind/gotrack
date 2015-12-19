package dhcp

import (
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
)

type Provider interface {
	Leases() ([]Lease, error)
}

func NewProvider(config *config.DhcpConfig, parent logging.Logger) (Provider, error) {
	switch config.Type {
	case "dnsmasq":
		return newDnsmasqProvider(config, parent)
	default:
		return nil, errors.Errorf("Unknown provider type: %s", config.Type)
	}
}
