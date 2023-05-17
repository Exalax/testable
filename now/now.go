package now

import "time"

var now time.Time

func Now() time.Time {
	if !now.IsZero() {
		return now
	}
	return time.Now()
}

func Set(t time.Time) {
	now = t
}

func Reset() {
	now = time.Time{}
}
