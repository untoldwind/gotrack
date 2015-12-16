package store

import (
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/conntrack"
	"github.com/untoldwind/gotrack/server/logging"
	"sync"
	"time"
)

type memoryStore struct {
	lock       sync.RWMutex
	provider   conntrack.Provider
	totals5Min *rrd
	ticker     *time.Ticker
	logger     logging.Logger
}

func newMemoryStore(config *config.StoreConfig, provider conntrack.Provider, parent logging.Logger) (*memoryStore, error) {
	store := &memoryStore{
		provider:   provider,
		totals5Min: newRRD(300, 1),
		ticker:     time.NewTicker(1 * time.Second),
		logger:     parent.WithContext(map[string]interface{}{"package": "store"}),
	}

	go store.pollData()

	return store, nil
}

func (s *memoryStore) Totals() *Span {
	return s.totals5Min.getSpan()
}

func (s *memoryStore) Stop() {
	s.ticker.Stop()
}

func (s *memoryStore) pollData() {
	for _ = range s.ticker.C {
		if totals, err := s.provider.Totals(); err == nil {
			s.logger.Debugf("Updating totals: %v", totals)
			s.totals5Min.addTotals(&totals.Receive, &totals.Send)
		} else {
			s.logger.Warn("Failed to get totals: %v", err)
		}
	}
}
