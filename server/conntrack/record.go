package conntrack

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/go-errors/errors"
)

type Connection struct {
	Transfer
	SrcIp   string `json:"src_ip"`
	SrcPort uint16 `json:"src_port"`
	DstIp   string `json:"dst_ip"`
	DstPort uint16 `json:"dst_port"`
}

// Conntrack record
// Resembles a single tracked connection
type Record struct {
	Protocol string     `json:"protocol"`
	Ttl      uint64     `json:"ttl"`
	State    string     `json:"state"`
	Send     Connection `json:"send"`
	Receive  Connection `json:"receive"`
	Key      string     `json:"key"`
}

func parseRecord(line string) (*Record, error) {
	fields := strings.Fields(line)

	if len(fields) < 12 {
		return nil, errors.Errorf("Invalid conntrack recod: %s", line)
	}

	hash := sha256.New()
	protocol := fields[0]

	hash.Write([]byte(protocol))

	ttl, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	var state string
	if strings.IndexRune(fields[5], '=') < 0 {
		state = fields[5]
		fields = fields[6:]
	} else {
		fields = fields[5:]
	}

	sendConnection, fields, err := parseConnection(fields)
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 2)
	hash.Write([]byte(sendConnection.SrcIp))
	binary.LittleEndian.PutUint16(buffer, sendConnection.SrcPort)
	hash.Write(buffer)
	hash.Write([]byte(sendConnection.DstIp))
	binary.LittleEndian.PutUint16(buffer, sendConnection.DstPort)
	hash.Write(buffer)

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
		Key:      base64.RawStdEncoding.EncodeToString(hash.Sum(nil)),
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
			result.SrcIp = field[sep+1:]
		case "dst":
			if idx >= 4 {
				return result, fields[idx:], nil
			}
			result.DstIp = field[sep+1:]
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
