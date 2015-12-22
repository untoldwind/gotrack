package store

type Device struct {
	Name            string `json:"name"`
	MacAddress      string `json:"mac_address"`
	IpAddress       string `json:"ip_address"`
	ConnectionCount int    `json:"connection_count"`
}
