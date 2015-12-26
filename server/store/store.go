package store

import (
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/conntrack"
	"github.com/untoldwind/gotrack/server/dhcp"
	"github.com/untoldwind/gotrack/server/logging"
)

type Store interface {
	Devices() []*Device
	DeviceDetails(deviceIp string) *DeviceDetails
	DeviceSpan(deviceIp string) *Span
	TotalsSpan() *Span
	TotalsRates() *Rates
	Stop()
}

func NewStore(config *config.StoreConfig, conntrackProvider conntrack.Provider, dhcpProvicer dhcp.Provider, parent logging.Logger) (Store, error) {
	switch config.Type {
	case "memory":
		return newMemoryStore(config, conntrackProvider, dhcpProvicer, parent)
	}
	return nil, errors.Errorf("Unknown story type: %s", config.Type)
}
