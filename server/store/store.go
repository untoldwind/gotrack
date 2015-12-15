package store

import (
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
)

type Store interface {
}

func NewStore(config *config.StoreConfig, parent logging.Logger) (*Store, error) {
	return nil, nil
}
