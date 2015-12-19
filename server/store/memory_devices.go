package store

import (
	"github.com/untoldwind/gotrack/server/dhcp"
	"sync"
)

type memoryDevices struct {
	lock    sync.RWMutex
	devices map[string]*Device
}

func newMemoryDevices() *memoryDevices {
	return &memoryDevices{
		devices: make(map[string]*Device, 0),
	}
}
func (d *memoryDevices) update(leases []dhcp.Lease) {
	d.lock.Lock()
	defer d.lock.Unlock()

	for _, lease := range leases {
		if device, ok := d.devices[lease.MacAddress]; ok {
			device.Name = lease.Name
			device.IpAddress = lease.IpAddress
		} else {
			d.devices[lease.MacAddress] = &Device{
				Name:       lease.Name,
				MacAddress: lease.MacAddress,
				IpAddress:  lease.IpAddress,
			}
		}
	}
}

func (d *memoryDevices) getDevices() []*Device {
	d.lock.RLock()
	defer d.lock.RUnlock()

	result := make([]*Device, 0, len(d.devices))
	for _, device := range d.devices {
		result = append(result, device)
	}
	return result
}
