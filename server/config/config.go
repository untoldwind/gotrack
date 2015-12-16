package config

import (
	"github.com/untoldwind/gotrack/server/logging"
	"os"
	"path"
	"path/filepath"
)

type Config struct {
	Server   *ServerConfig   `json:"server" yaml:"server"`
	Provider *ProviderConfig `json:"provider" yaml:"provider"`
	Store    *StoreConfig    `json:"store" yaml:"store"`
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

	for _, name := range []string{"config.json", "config.yaml"} {
		fileName := path.Join(absoluteConfigDir, name)
		if _, err := os.Stat(fileName); err == nil {
			if err := loadConfigFile(fileName, &config); err == nil {
				return &config, nil
			} else {
				return nil, err
			}
		}
	}

	logger.Warn("Read config failed (will use defaults)")
	return &config, nil
}
