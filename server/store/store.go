package store

import (
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/conntrack"
	"github.com/untoldwind/gotrack/server/logging"
)

type Store interface {
	Totals() *Span
	Stop()
}

func NewStore(config *config.StoreConfig, provider conntrack.Provider, parent logging.Logger) (Store, error) {
	switch config.Type {
	case "memory":
		return newMemoryStore(config, provider, parent)
	}
	return nil, errors.Errorf("Unknown story type: %s", config.Type)
}
