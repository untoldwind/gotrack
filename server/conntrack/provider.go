package conntrack

import (
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
)

type Provider interface {
	Records() ([]*Record, error)
}

func NewProvider(config *config.ProviderConfig, parent logging.Logger) (Provider, error) {
	switch config.Type {
	case "proc":
		return newProcProvider(config, parent)
	default:
		return nil, errors.Errorf("Unknown provider type: %s", config.Type)
	}
}
