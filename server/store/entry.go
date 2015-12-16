package store

type Entry struct {
	BytesIn    uint64 `json:"bytes_in"`
	PacketsIn  uint64 `json:"packets_in"`
	BytesOut   uint64 `json:"bytes_out"`
	PacketsOut uint64 `json:"packets_out"`
}
