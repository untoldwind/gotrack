package store

import "time"

type Span struct {
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	CurrentIdx int       `json:"current_idx"`
	Max        Entry     `json:"max"`
	Deltas     []Entry   `json:"deltas"`
}
