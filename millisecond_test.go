package timestamp

import (
	"testing"
	"time"
)

func TestMillisecond(t *testing.T) {
	tm := time.Date(2019, 12, 31, 1, 29, 30, 123456789, time.UTC)
	s := FromTimeToMillisecond(tm)

	if got, want := s, Millisecond(1577755770123); got != want {
		t.Errorf("FromTimeToMillisecond result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.ToTime(), tm.Truncate(time.Millisecond); !got.Equal(want) {
		t.Errorf("ToTime result unmatch, got=%+v, want=%+v", got, want)
	}

	d := 3 * time.Millisecond
	if got, want := s.Add(d), Millisecond(1577755770126); got != want {
		t.Errorf("Add result unmatch, got=%d, want=%d", got, want)
	}
	if got, want := s.Sub(Millisecond(1577755770120)), d; got != want {
		t.Errorf("Sub result unmatch, got=%s, want=%s", got, want)
	}
}
