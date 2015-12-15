package conntrack

import (
	"bufio"
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"os"
	"strings"
	"unicode"
)

type procProvider struct {
	config *config.ProviderConfig
	logger logging.Logger
}

func newProcProvider(config *config.ProviderConfig, parent logging.Logger) (*procProvider, error) {
	return &procProvider{
		config: config,
		logger: parent.WithContext(map[string]interface{}{"package": "conntrack"}),
	}, nil
}

func (c *procProvider) Totals() (*Totals, error) {
	file, err := os.Open(c.config.DevFile)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimLeftFunc(scanner.Text(), unicode.IsSpace)
		if strings.HasPrefix(line, c.config.WanInterface) {
			return parseTotals(line)
		}
	}
	return nil, errors.Errorf("No stats for interface found: %s", c.config.WanInterface)
}

func (c *procProvider) Records() ([]*Record, error) {
	file, err := os.Open(c.config.ConntrackFile)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	defer file.Close()

	result := make([]*Record, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if record, err := parseRecord(scanner.Text()); err == nil {
			result = append(result, record)
		} else {
			c.logger.Warnf("Parse error: %v", err)
		}
	}

	return result, nil
}
