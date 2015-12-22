package store

type Entry struct {
	BytesIn    uint64 `json:"bytes_in"`
	PacketsIn  uint64 `json:"packets_in"`
	BytesOut   uint64 `json:"bytes_out"`
	PacketsOut uint64 `json:"packets_out"`
}

func (e Entry) delta(other Entry) Entry {
	return Entry{
		BytesIn:    uintDelta(e.BytesIn, other.BytesIn),
		PacketsIn:  uintDelta(e.PacketsIn, other.PacketsIn),
		BytesOut:   uintDelta(e.BytesOut, other.BytesOut),
		PacketsOut: uintDelta(e.PacketsOut, other.PacketsOut),
	}
}

func (e *Entry) maxOf(other Entry) {
	e.BytesIn = uintMax(e.BytesIn, other.BytesIn)
	e.PacketsIn = uintMax(e.PacketsIn, other.PacketsIn)
	e.BytesOut = uintMax(e.BytesOut, other.BytesOut)
	e.PacketsOut = uintMax(e.PacketsOut, other.PacketsOut)
}

func (e *Entry) add(entry Entry) {
	e.BytesIn += entry.BytesIn
	e.PacketsIn += entry.PacketsIn
	e.BytesOut += entry.BytesOut
	e.PacketsOut += entry.PacketsOut
}

func (e *Entry) div(div uint64) {
	e.BytesIn /= div
	e.PacketsIn /= div
	e.BytesOut /= div
	e.PacketsOut /= div
}

func uintDelta(v1, v2 uint64) uint64 {
	if v2 >= v1 {
		return v2 - v1
	}
	return 0
}

func uintMax(v1, v2 uint64) uint64 {
	if v2 > v1 {
		return v2
	}
	return v1
}
