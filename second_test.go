package timestamp

import (
	"testing"
	"time"
)

func TestSecond(t *testing.T) {
	tm := time.Date(2019, 12, 31, 1, 29, 30, 123456789, time.UTC)
	s := FromTimeToSecond(tm)

	if got, want := s, Second(1577755770); got != want {
		t.Errorf("FromTimeToSecond result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.ToTime(), tm.Truncate(time.Second); !got.Equal(want) {
		t.Errorf("ToTime result unmatch, got=%+v, want=%+v", got, want)
	}

	d := 3 * time.Second
	if got, want := s.Add(d), Second(1577755773); got != want {
		t.Errorf("Add result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.Sub(Second(1577755767)), d; got != want {
		t.Errorf("Sub result unmatch, got=%s, want=%s", got, want)
	}
}
