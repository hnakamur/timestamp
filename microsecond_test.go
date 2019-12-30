package timestamp

import (
	"testing"
	"time"
)

func TestMicrosecond(t *testing.T) {
	tm := time.Date(2019, 12, 31, 1, 29, 30, 123456789, time.UTC)
	s := FromTimeToMicrosecond(tm)

	if got, want := s, Microsecond(1577755770123456); got != want {
		t.Errorf("FromTimeToMicrosecond result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.ToTime(), tm.Truncate(time.Microsecond); !got.Equal(want) {
		t.Errorf("ToTime result unmatch, got=%+v, want=%+v", got, want)
	}

	d := 3 * time.Microsecond
	if got, want := s.Add(d), Microsecond(1577755770123459); got != want {
		t.Errorf("Add result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.Sub(Microsecond(1577755770123453)), d; got != want {
		t.Errorf("Sub result unmatch, got=%s, want=%s", got, want)
	}
}
