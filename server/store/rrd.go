package store

import (
	"github.com/untoldwind/gotrack/server/conntrack"
	"sync"
	"time"
)

var rrd_now = func() int64 {
	return time.Now().Unix()
}

type rrd struct {
	lock    sync.RWMutex
	entries []Entry
	start   int64
	end     int64
	step    int64
	size    int64
}

func newRRD(span, step int64) *rrd {
	now := rrd_now()
	now -= now % step

	return &rrd{
		entries: make([]Entry, span/step),
		start:   now,
		end:     now + span,
		step:    step,
		size:    span / step,
	}
}

func (r *rrd) addTotals(in, out *conntrack.Transfer) {
	r.lock.Lock()
	defer r.lock.Unlock()

	now := rrd_now()
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
}

func (r *rrd) getSpan() *Span {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result := &Span{
		Start:   time.Unix(r.start, 0),
		End:     time.Unix(r.end, 0),
		Entries: make([]Entry, len(r.entries)),
	}
	idx := (r.start / r.step) % r.size

	len := copy(result.Entries, r.entries[idx:])
	if idx > 0 {
		copy(result.Entries[len:], r.entries[:idx])
	}
	return result
}
