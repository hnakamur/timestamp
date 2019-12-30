package timestamp

import (
	"testing"
	"time"
)

func TestNanosecond(t *testing.T) {
	tm := time.Date(2019, 12, 31, 1, 29, 30, 123456789, time.UTC)
	s := FromTimeToNanosecond(tm)

	if got, want := s, Nanosecond(1577755770123456789); got != want {
		t.Errorf("FromTimeToNanosecond result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.ToTime(), tm.Truncate(time.Nanosecond); !got.Equal(want) {
		t.Errorf("ToTime result unmatch, got=%+v, want=%+v", got, want)
	}

	d := 3 * time.Nanosecond
	if got, want := s.Add(d), Nanosecond(1577755770123456792); got != want {
		t.Errorf("Add result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.Sub(Nanosecond(1577755770123456786)), d; got != want {
		t.Errorf("Sub result unmatch, got=%s, want=%s", got, want)
	}
}
