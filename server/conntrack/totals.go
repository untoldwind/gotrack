package conntrack

import (
	"github.com/go-errors/errors"
	"strconv"
	"strings"
)

type Totals struct {
	Send    Transfer `json:"send"`
	Receive Transfer `json:"receive"`
}

func parseTotals(line string) (*Totals, error) {
	fields := strings.Fields(line)

	if len(fields) < 17 {
		return nil, errors.Errorf("Invalid interface stats: %s", line)
	}

	receivedBytes, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	receivedPackets, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	sendBytes, err := strconv.ParseUint(fields[9], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	sendPackets, err := strconv.ParseUint(fields[10], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return &Totals{
		Send: Transfer{
			Bytes:   sendBytes,
			Packets: sendPackets,
		},
		Receive: Transfer{
			Bytes:   receivedBytes,
			Packets: receivedPackets,
		},
	}, nil
}
