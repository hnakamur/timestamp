package timestamp

import "time"

// Millisecond is the Unix timestamp, the number of milliseconds
// elapsed since January 1, 1970 UTC.
type Millisecond int64

// FromTimeToMillisecond returns t as a Millisecond.
func FromTimeToMillisecond(t time.Time) Millisecond {
	return Millisecond(t.UnixNano() / 1e6)
}

// ToTime returns m as a time.Time.
func (m Millisecond) ToTime() time.Time {
	return time.Unix(int64(m)/1e3, (int64(m)%1e3)*1e6)
}

// Add returns m+d.
func (m Millisecond) Add(d time.Duration) Millisecond {
	return m + Millisecond(d/time.Millisecond)
}

// Sub returns the duration m-t.
// To compute m-d for duration, use m.Add(-d).
func (m Millisecond) Sub(t Millisecond) time.Duration {
	return time.Duration(m-t) * time.Millisecond
}
