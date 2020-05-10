package edsm

import (
	"fmt"
	"strings"
	"time"
)

const ctLayout = "2006-01-02 15:04:05"

func formateDateTime(value time.Time) string {
	return value.UTC().Format(ctLayout)
}

// Time EDSM formatted time
type Time struct {
	time.Time
}

// UnmarshalJSON deserializes from JSON
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse(ctLayout, s)
	return
}

// MarshalJSON serializes to JSON
func (t *Time) MarshalJSON() ([]byte, error) {
	if t.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

// IsSet return whether this time is not null
func (t *Time) IsSet() bool {
	return t.UnixNano() != nilTime
}
