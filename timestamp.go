package msg

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

// NewTimestamp creates a new Timestamp for u
func NewTimestamp(u time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: u.Unix(),
		Nanos:   int32(u.Nanosecond()),
	}
}

// Time returns the Time corresponding to the Timestamp
func Time(m *timestamp.Timestamp) time.Time {
	return time.Unix(m.Seconds, int64(m.Nanos))
}

// TimestampEqual returns if one Timestamp is equal to the other
func TimestampEqual(m, n *timestamp.Timestamp) bool {
	return m.Seconds == n.Seconds && m.Nanos == n.Nanos
}
