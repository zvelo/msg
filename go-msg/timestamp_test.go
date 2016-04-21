package msg

import (
	"testing"
	"time"
)

func TestBeginning(t *testing.T) {
	// unix time starts on Jan 1, 1970 UTC
	u := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	ts := NewTimestamp(u)

	if ts == nil {
		t.Error("ts should not be nil")
	}

	if ts.Seconds != 0 {
		t.Error("ts.Seconds should equal 0")
	}

	if ts.Nanos != 0 {
		t.Error("ts.Nanos should equal 0")
	}

	if !u.Equal(Time(ts)) {
		t.Error("u.Equal(Time(ts) should be true")
	}
}

func TestBeforeBeginning(t *testing.T) {
	// Israel is UTC+2 (7200s)
	loc, err := time.LoadLocation("Israel")
	if err != nil {
		t.Error("err should be nil")
	}

	u := time.Date(1970, time.January, 1, 0, 0, 0, 0, loc)
	ts := NewTimestamp(u)

	if ts == nil {
		t.Error("ts should not be nil")
	}

	if ts.Seconds != -7200 {
		t.Error("ts.Seconds should equal -7200")
	}

	if ts.Nanos != 0 {
		t.Error("ts.Nanos should equal 0")
	}

	if !u.Equal(Time(ts)) {
		t.Error("u.Equal(Time(ts) should be true")
	}
}

func TestAfterBeginning(t *testing.T) {
	// Denver is UTC-7 (25200s)
	loc, err := time.LoadLocation("America/Denver")
	if err != nil {
		t.Error("err should be nil")
	}

	u := time.Date(1970, time.January, 1, 0, 0, 0, 0, loc)
	ts := NewTimestamp(u)

	if ts == nil {
		t.Error("ts should not be nil")
	}

	if ts.Seconds != 25200 {
		t.Error("ts.Seconds should equal 25200")
	}

	if ts.Nanos != 0 {
		t.Error("ts.Nanos should equal 0")
	}

	if !u.Equal(Time(ts)) {
		t.Error("u.Equal(Time(ts) should be true")
	}
}

func TestWithNano(t *testing.T) {
	u := time.Date(1970, time.January, 1, 0, 0, 0, 999999999, time.UTC)
	ts := NewTimestamp(u)

	if ts == nil {
		t.Error("ts should not be nil")
	}

	if ts.Seconds != 0 {
		t.Error("ts.Seconds should equal 0")
	}

	if ts.Nanos != 999999999 {
		t.Error("ts.Nanos should equal 999999999")
	}

	if !u.Equal(Time(ts)) {
		t.Error("u.Equal(Time(ts) should be true")
	}
}

func TestNow(t *testing.T) {
	u := time.Now()
	ts := NewTimestamp(u)

	if ts == nil {
		t.Error("ts should not be nil")
	}

	if ts.Seconds <= 1426114544 {
		t.Error("ts.Seconds should be less than or equal to 1426114544")
	}

	if !u.Equal(Time(ts)) {
		t.Error("u.Equal(Time(ts) should be true")
	}
}
