package now

import "time"

var now time.Time

func Now() time.Time {
	if !now.IsZero() {
		return now
	}
	return time.Now()
}

func Until(t time.Time) time.Duration {
	return t.Sub(Now())
}

func Set(t time.Time) {
	now = t
}

func Reset() {
	now = time.Time{}
}
