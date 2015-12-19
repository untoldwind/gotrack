package dhcp

type Lease struct {
	Name       string `json:"name"`
	IpAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
	ExpiresAt  int64  `json:"expires_at"`
}
