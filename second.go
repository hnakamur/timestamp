package timestamp

import "time"

// Second is the Unix timestamp, the number of seconds
// elapsed since January 1, 1970 UTC.
type Second int64

// FromTimeToSecond returns t as a Second.
func FromTimeToSecond(t time.Time) Second {
	return Second(t.Unix())
}

// ToTime returns s as a time.Time.
func (s Second) ToTime() time.Time {
	return time.Unix(int64(s), 0)
}

// Add returns s+d.
func (s Second) Add(d time.Duration) Second {
	return s + Second(d/time.Second)
}

// Sub returns the duration s-t.
// To compute s-d for duration, use s.Add(-d).
func (s Second) Sub(t Second) time.Duration {
	return time.Duration(s-t) * time.Second
}
