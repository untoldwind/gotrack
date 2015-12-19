package store

import (
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/conntrack"
	"github.com/untoldwind/gotrack/server/dhcp"
	"github.com/untoldwind/gotrack/server/logging"
	"time"
)

type memoryStore struct {
	conntrackProvider conntrack.Provider
	dhcpProvider      dhcp.Provider
	devices           *memoryDevices
	totals5Min        *rrd
	dhcpTicker        *time.Ticker
	totalsTicker      *time.Ticker
	logger            logging.Logger
}

func newMemoryStore(config *config.StoreConfig, conntrackProvider conntrack.Provider, dhcpProvider dhcp.Provider, parent logging.Logger) (*memoryStore, error) {
	store := &memoryStore{
		conntrackProvider: conntrackProvider,
		dhcpProvider:      dhcpProvider,
		devices:           newMemoryDevices(),
		totals5Min:        newRRD(time.Now(), 300, 1),
		dhcpTicker:        time.NewTicker(10 * time.Second),
		totalsTicker:      time.NewTicker(1 * time.Second),
		logger:            parent.WithContext(map[string]interface{}{"package": "store"}),
	}

	go store.pollDhcp()
	go store.pollTotals()

	return store, nil
}

func (s *memoryStore) Devices() []*Device {
	return s.devices.getDevices()
}

func (s *memoryStore) TotalsSpan() *Span {
	return s.totals5Min.getSpan()
}

func (s *memoryStore) TotalsRates() *Rates {
	return &Rates{
		Rate5Sec: s.totals5Min.getRate(5),
	}
}

func (s *memoryStore) Stop() {
	s.dhcpTicker.Stop()
	s.totalsTicker.Stop()
}

func (s *memoryStore) pollDhcp() {
	for _ = range s.dhcpTicker.C {
		if leases, err := s.dhcpProvider.Leases(); err == nil {
			s.logger.Debugf("Updating leases: %v", leases)
			s.devices.update(leases)
		} else {
			s.logger.Warn("Failed to get leases: %v", err)
		}
	}
}

func (s *memoryStore) pollTotals() {
	for time := range s.totalsTicker.C {
		if totals, err := s.conntrackProvider.Totals(); err == nil {
			s.logger.Debugf("Updating totals: %v", totals)
			s.totals5Min.addTotals(time, &totals.Receive, &totals.Send)
		} else {
			s.logger.Warn("Failed to get totals: %v", err)
		}
	}
}
