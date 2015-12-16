package config

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func isConfigFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".json") || strings.HasSuffix(fileName, ".yaml")
}

func loadConfigFile(fileName string, configValue interface{}) error {
	if !isConfigFile(fileName) {
		return errors.Errorf("%s is not a configuration file", fileName)
	}

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(fileName, ".json"):
		err := json.Unmarshal(data, configValue)
		if err != nil {
			return errors.Errorf("JSON error in %s: %s", fileName, err.Error())
		}
		return nil
	case strings.HasSuffix(fileName, ".yaml"):
		err := yaml.Unmarshal(data, configValue)
		if err != nil {
			return errors.Errorf("JSON error in %s: %s", fileName, err.Error())
		}
		return nil
	default:
		return errors.Errorf("%s is not a valid config file", fileName)
	}
}
