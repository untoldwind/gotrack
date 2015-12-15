package config

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/untoldwind/gotrack/server/logging"
)

type Config struct {
	Server   *ServerConfig
	Provider *ProviderConfig
	Store    *StoreConfig
}

func NewConfig(configDir string, logger logging.Logger) (*Config, error) {
	absoluteConfigDir, err := filepath.Abs(configDir)
	if err != nil {
		return nil, err
	}

	config := Config{
		Server:   newServerConfig(),
		Provider: newProviderConfig(),
		Store:    newStoreConfig(),
	}
	files, err := ioutil.ReadDir(absoluteConfigDir)
	if err != nil {
		logger.Warnf("Read config failed (will use defaults): %s", err.Error())
		return &config, nil
	}
	for _, file := range files {
		switch {
		case !file.IsDir() && strings.HasPrefix(file.Name(), "server."):
			var err error
			config.Server, err = readServerConfig(path.Join(absoluteConfigDir, file.Name()))
			if err != nil {
				return nil, err
			}
		case !file.IsDir() && strings.HasPrefix(file.Name(), "provider."):
			var err error
			config.Provider, err = readProviderConfig(path.Join(absoluteConfigDir, file.Name()))
			if err != nil {
				return nil, err
			}
		case !file.IsDir() && strings.HasPrefix(file.Name(), "store."):
			var err error
			config.Store, err = readStoreConfig(path.Join(absoluteConfigDir, file.Name()))
			if err != nil {
				return nil, err
			}
		}
	}

	return &config, nil
}
