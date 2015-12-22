package store

type DeviceDetails struct {
	Device
	Connections []*Connection `json:"connections"`
}
