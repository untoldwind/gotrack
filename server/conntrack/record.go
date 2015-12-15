package conntrack

import (
	"github.com/go-errors/errors"
	"strconv"
	"strings"
)

type Connection struct {
	Src     string `json:"src"`
	SrcPort uint16 `json:"src_port"`
	Dst     string `json:"destination"`
	DstPort uint16 `json:"dst_port"`
	Packets uint64 `json:"packets"`
	Bytes   uint64 `json:"bytes"`
}

// Conntrack record
// Resembles a single tracked connection
type Record struct {
	Protocol string     `json:"protocol"`
	Ttl      uint64     `json:"ttl"`
	State    string     `json:"state"`
	Send     Connection `json:"send"`
	Receive  Connection `json:"receive"`
}

func parseRecord(line string) (*Record, error) {
	fields := strings.Fields(line)

	if len(fields) < 12 {
		return nil, errors.Errorf("Invalid conntrack recod: %s", line)
	}

	protocol := fields[0]

	ttl, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	var state string
	if strings.IndexRune(fields[3], '=') < 0 {
		state = fields[3]
		fields = fields[4:]
	} else {
		fields = fields[3:]
	}

	sendConnection, fields, err := parseConnection(fields)
	if err != nil {
		return nil, err
	}

	receiveConnection, _, err := parseConnection(fields)
	if err != nil {
		return nil, err
	}

	result := &Record{
		Protocol: protocol,
		Ttl:      ttl,
		State:    state,
		Send:     *sendConnection,
		Receive:  *receiveConnection,
	}
	return result, nil
}

func parseConnection(fields []string) (*Connection, []string, error) {
	if len(fields) < 4 {
		return nil, nil, errors.Errorf("Not enough fields: %v", fields)
	}
	result := &Connection{}

	for idx, field := range fields {
		sep := strings.IndexRune(field, '=')
		if sep < 0 || idx >= 6 {
			return result, fields[idx:], nil
		}
		switch field[:sep] {
		case "src":
			if idx >= 4 {
				return result, fields[idx:], nil
			}
			result.Src = field[sep+1:]
		case "dst":
			if idx >= 4 {
				return result, fields[idx:], nil
			}
			result.Dst = field[sep+1:]
		case "sport":
			if idx >= 4 {
				return result, fields[idx:], nil
			}
			if port, err := strconv.ParseUint(field[sep+1:], 10, 16); err == nil {
				result.SrcPort = uint16(port)
			} else {
				return nil, nil, errors.Wrap(err, 0)
			}
		case "dport":
			if idx >= 4 {
				return result, fields[idx:], nil
			}
			if port, err := strconv.ParseUint(field[sep+1:], 10, 16); err == nil {
				result.DstPort = uint16(port)
			} else {
				return nil, nil, errors.Wrap(err, 0)
			}
		case "packets":
			if packets, err := strconv.ParseUint(field[sep+1:], 10, 64); err == nil {
				result.Packets = packets
			} else {
				return nil, nil, errors.Wrap(err, 0)
			}
		case "bytes":
			if bytes, err := strconv.ParseUint(field[sep+1:], 10, 64); err == nil {
				result.Bytes = bytes
			} else {
				return nil, nil, errors.Wrap(err, 0)
			}
		}
	}

	return result, nil, nil
}
