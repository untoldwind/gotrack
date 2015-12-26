package store

type Connection struct {
	Protocol string `json:"protocol"`
	SrcPort  uint16 `json:"src_port"`
	DestHost string `json:"dst_host"`
	DestPort uint16 `json:"dst_port"`
	Totals   Entry  `json:"totals"`
}
