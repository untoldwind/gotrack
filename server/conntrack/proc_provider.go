package conntrack

import (
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"os"
	"bufio"
)

type procProvider struct {
	fileName string
	logger   logging.Logger
}

func newProcProvider(config *config.ProviderConfig, parent logging.Logger) (*procProvider, error) {
	return &procProvider{
		fileName: config.ProcFile,
		logger:   parent.WithContext(map[string]interface{}{"package": "conntrack"}),
	}, nil
}

func (c *procProvider) Records() ([]*Record, error) {
	file, err := os.Open(c.fileName)
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
