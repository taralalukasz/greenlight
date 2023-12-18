package data

import (
	"time"
)

// annotations on model are instructions for json.Marshall to change the names of the fields
// more information can be found here https://pkg.go.dev/encoding/json
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // ignore on serialization
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`           // omitempty means it will be ignored is nil / "" / {}
	Runtime   int32     `json:"runtime,omitempty,string"` // we can force to represent int as string
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
