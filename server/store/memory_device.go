package store

import (
	"github.com/untoldwind/gotrack/server/conntrack"
	"time"
)

type memoryDevice struct {
	name        string
	ipAddress   string
	macAddress  string
	totals5Min  *rrd
	totals      Entry
	baseTotals  Entry
	connections map[string]*Connection
}

func newMemoryDevice(name, ipAddress, macAddress string) *memoryDevice {
	return &memoryDevice{
		name:        name,
		ipAddress:   ipAddress,
		macAddress:  macAddress,
		totals5Min:  newRRD(time.Now(), 300, 1),
		connections: make(map[string]*Connection, 0),
	}
}

func (d *memoryDevice) updateConnections(record *conntrack.Record) {
	connection, ok := d.connections[record.Key]

	if !ok {
		connection = &Connection{
			Protocol: record.Protocol,
			SrcPort:  record.Send.SrcPort,
			DestHost: record.Send.DstIp,
			DestPort: record.Send.DstPort,
		}
		d.connections[record.Key] = connection
	}
	connection.Totals.BytesIn = record.Receive.Bytes
	connection.Totals.PacketsIn = record.Receive.Packets
	connection.Totals.BytesOut = record.Send.Bytes
	connection.Totals.PacketsOut = record.Send.Packets
}

func (d *memoryDevice) cleanupConnections(time time.Time, activeConnections map[string]bool) {
	d.totals = d.baseTotals
	for key, connection := range d.connections {
		d.totals.add(connection.Totals)
		if !activeConnections[key] {
			d.baseTotals.add(connection.Totals)
			delete(d.connections, key)
		}
	}
	d.totals5Min.addEntry(time, d.totals)
}

func (d *memoryDevice) toDevice() *Device {
	return &Device{
		Name:            d.name,
		IpAddress:       d.ipAddress,
		MacAddress:      d.macAddress,
		ConnectionCount: len(d.connections),
		Totals:          d.totals,
		Rate5Sec:        d.totals5Min.getRate(5),
	}
}

func (d *memoryDevice) toDeviceDetails() *DeviceDetails {
	connections := make([]*Connection, 0, len(d.connections))

	for _, connection := range d.connections {
		connections = append(connections, connection)
	}
	return &DeviceDetails{
		Device:      *d.toDevice(),
		Connections: connections,
	}
}

func (d *memoryDevice) getSpan() *Span {
	return d.totals5Min.getSpan()
}
