package store

import (
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/conntrack"
	"github.com/untoldwind/gotrack/server/dhcp"
	"net"
	"sync"
)

type memoryDevices struct {
	lock       sync.RWMutex
	lanNetwork *net.IPNet
	devices    map[string]*memoryDevice
}

func newMemoryDevices(config *config.StoreConfig) (*memoryDevices, error) {
	_, lanNetwork, err := net.ParseCIDR(config.LanCIDR)

	if err != nil {
		return nil, err
	}

	return &memoryDevices{
		devices:    make(map[string]*memoryDevice, 0),
		lanNetwork: lanNetwork,
	}, nil
}

func (d *memoryDevices) updateConntrackRecords(records []*conntrack.Record) {
	d.lock.Lock()
	defer d.lock.Unlock()

	keys := make(map[string]bool, 0)
	for _, record := range records {
		keys[record.Key] = true
		if d.lanNetwork.Contains(net.ParseIP(record.Send.SrcIp)) &&
			!d.lanNetwork.Contains(net.ParseIP(record.Send.DstIp)) {

			device, ok := d.devices[record.Send.SrcIp]
			if !ok {
				device = newMemoryDevice(record.Send.SrcIp, record.Send.SrcIp, "")
				d.devices[record.Send.SrcIp] = device
			}
			device.updateConnections(record)
		}
	}
	for _, device := range d.devices {
		device.cleanupConnections(keys)
	}
}

func (d *memoryDevices) updateLeases(leases []dhcp.Lease) {
	d.lock.Lock()
	defer d.lock.Unlock()

	for _, lease := range leases {
		if device, ok := d.devices[lease.IpAddress]; ok {
			device.name = lease.Name
			device.ipAddress = lease.IpAddress
		} else {
			d.devices[lease.IpAddress] = newMemoryDevice(lease.Name, lease.IpAddress, lease.MacAddress)
		}
	}
}

func (d *memoryDevices) getDevices() []*Device {
	d.lock.RLock()
	defer d.lock.RUnlock()

	result := make([]*Device, 0, len(d.devices))
	for _, device := range d.devices {
		result = append(result, device.toDevice())
	}
	return result
}

func (d *memoryDevices) getDeviceDetails(deviceIp string) *DeviceDetails {
	d.lock.RLock()
	defer d.lock.RUnlock()

	if device, ok := d.devices[deviceIp]; ok {
		return device.toDeviceDetails()
	}
	return nil
}
