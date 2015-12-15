package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func isConfigFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".json") || strings.HasSuffix(fileName, ".yaml")
}

func loadConfigFile(fileName string, configValue interface{}) error {
	if !isConfigFile(fileName) {
		return fmt.Errorf("%s is not a configuration file", fileName)
	}

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(fileName, ".json"):
		err := json.Unmarshal(data, configValue)
		if err != nil {
			return fmt.Errorf("JSON error in %s: %s", fileName, err.Error())
		}
		return nil
	case strings.HasSuffix(fileName, ".yaml"):
		err := yaml.Unmarshal(data, configValue)
		if err != nil {
			return fmt.Errorf("JSON error in %s: %s", fileName, err.Error())
		}
		return nil
	default:
		return fmt.Errorf("%s is not a valid config file", fileName)
	}
}
