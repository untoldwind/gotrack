package conntrack

type Transfer struct {
	Packets uint64 `json:"packets"`
	Bytes   uint64 `json:"bytes"`
}
