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
	conntrackTicker   *time.Ticker
	dhcpTicker        *time.Ticker
	totalsTicker      *time.Ticker
	logger            logging.Logger
}

func newMemoryStore(config *config.StoreConfig,
	conntrackProvider conntrack.Provider, dhcpProvider dhcp.Provider, parent logging.Logger) (*memoryStore, error) {

	devices, err := newMemoryDevices(config)
	if err != nil {
		return nil, err
	}
	store := &memoryStore{
		conntrackProvider: conntrackProvider,
		dhcpProvider:      dhcpProvider,
		devices:           devices,
		totals5Min:        newRRD(time.Now(), 300, 1),
		conntrackTicker:   time.NewTicker(1 * time.Second),
		dhcpTicker:        time.NewTicker(10 * time.Second),
		totalsTicker:      time.NewTicker(1 * time.Second),
		logger:            parent.WithContext(map[string]interface{}{"package": "store"}),
	}

	go store.pollConntrack()
	go store.pollDhcp()
	go store.pollTotals()

	return store, nil
}

func (s *memoryStore) Devices() []*Device {
	return s.devices.getDevices()
}

func (s *memoryStore) DeviceDetails(deviceIp string) *DeviceDetails {
	return s.devices.getDeviceDetails(deviceIp)
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
	s.conntrackTicker.Stop()
	s.dhcpTicker.Stop()
	s.totalsTicker.Stop()
}

func (s *memoryStore) pollConntrack() {
	for time := range s.conntrackTicker.C {
		if records, err := s.conntrackProvider.Records(); err == nil {
			s.logger.Debugf("Update conntrack records: %v", records)
			s.devices.updateConntrackRecords(time, records)
		} else {
			s.logger.Warn("Failed to get conntrack records: %v", err)
		}
	}
}

func (s *memoryStore) pollDhcp() {
	for _ = range s.dhcpTicker.C {
		if leases, err := s.dhcpProvider.Leases(); err == nil {
			s.logger.Debugf("Updating leases: %v", leases)
			s.devices.updateLeases(leases)
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
