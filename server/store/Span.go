package store

import "time"

type Span struct {
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Entries []Entry   `json:"entry"`
}
