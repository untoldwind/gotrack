package store

type Connection struct {
	Protocol string `json:"protocol"`
	SrcIp    uint16 `json:"src_ip"`
	DestHost string `json:"dst_host"`
	DestIp   uint16 `json:"dst_ip"`
	Totals   Entry  `json:"totals"`
}
