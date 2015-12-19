package dhcp

import (
	"bufio"
	"github.com/go-errors/errors"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"os"
	"strconv"
	"strings"
)

type dnsmasqProvider struct {
	config *config.DhcpConfig
	logger logging.Logger
}

func newDnsmasqProvider(config *config.DhcpConfig, parent logging.Logger) (*dnsmasqProvider, error) {
	return &dnsmasqProvider{
		config: config,
		logger: parent.WithContext(map[string]interface{}{"package": "dhcp"}),
	}, nil
}

func (d *dnsmasqProvider) Leases() ([]Lease, error) {
	file, err := os.Open(d.config.DnsmasqFile)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	defer file.Close()

	result := make([]Lease, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) < 4 {
			d.logger.Warnf("Invalid dnsmsaq lease: %v", fields)
		}
		if expireAt, err := strconv.ParseInt(fields[0], 10, 64); err == nil {
			result = append(result, Lease{
				ExpiresAt:  expireAt,
				MacAddress: fields[1],
				IpAddress:  fields[2],
				Name:       fields[3],
			})
		} else {
			d.logger.Warnf("Invalid dnsmsaq lease: %v", fields)
		}
	}
	return result, nil
}
