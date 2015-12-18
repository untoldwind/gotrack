package store

import (
	"github.com/untoldwind/gotrack/server/conntrack"
	"math"
	"sync"
	"time"
)

type rrd struct {
	lock    sync.RWMutex
	entries []Entry
	start   int64
	end     int64
	last    int64
	step    int64
	size    int64
}

func newRRD(start time.Time, span, step int64) *rrd {
	startUnix := start.Unix()
	startUnix -= startUnix % step

	return &rrd{
		entries: make([]Entry, span/step),
		start:   startUnix,
		end:     startUnix + span,
		last:    math.MinInt64,
		step:    step,
		size:    span / step,
	}
}

func (r *rrd) addTotals(time time.Time, in, out *conntrack.Transfer) {
	r.lock.Lock()
	defer r.lock.Unlock()

	now := time.Unix()
	idx := (now / r.step) % r.size
	r.entries[idx] = Entry{
		BytesIn:    in.Bytes,
		PacketsIn:  in.Packets,
		BytesOut:   out.Bytes,
		PacketsOut: out.Packets,
	}
	if now >= r.end {
		now -= now % r.step
		r.end = now + r.step
		r.start = r.end - r.step*r.size
	}
	r.last = now
}

func (r *rrd) getSpan() *Span {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var currentIdx int
	if r.last >= r.start {
		currentIdx = int((r.last-r.start)/r.step - 1)
	}

	result := &Span{
		Start:      time.Unix(r.start, 0),
		End:        time.Unix(r.end, 0),
		Deltas:     make([]Entry, len(r.entries)-1),
		CurrentIdx: currentIdx,
	}
	startIdx := int((r.start / r.step) % r.size)

	for i := 0; i < len(r.entries)-1; i++ {
		idx1 := (startIdx + i) % len(r.entries)
		idx2 := (startIdx + i + 1) % len(r.entries)

		result.Deltas[i] = r.entries[idx1].delta(r.entries[idx2])
		result.Max.maxOf(result.Deltas[i])
	}

	return result
}

func (r *rrd) getRate(offset int64) Entry {
	var result Entry

	if offset > 0 && r.last >= r.start+offset {
		idx1 := (r.last / r.step) % r.size
		idx2 := ((r.last - offset) / r.step) % r.size
		result = r.entries[idx2].delta(r.entries[idx1])
		result.div(uint64(offset))
	}

	return result
}
