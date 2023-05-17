package now

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	Reset()

	assert.NotEqual(t, time.Time{}, Now())
}

func TestSet(t *testing.T) {
	someTime := time.Date(2015, time.May, 12, 18, 43, 21, 0, time.UTC)

	Set(someTime)
	defer Reset()

	assert.Equal(t, someTime, Now())
}

func TestReset(t *testing.T) {
	someTime := time.Date(2015, time.May, 12, 18, 43, 21, 0, time.UTC)

	Set(someTime)
	Reset()

	assert.NotEqual(t, someTime, Now())
}

func TestUntil(t *testing.T) {
	someTime := time.Date(2015, time.May, 12, 18, 43, 21, 0, time.UTC)

	Set(someTime)
	defer Reset()

	untilTime := time.Date(2015, time.May, 12, 20, 43, 21, 0, time.UTC)

	assert.Equal(t, 2*time.Hour, Until(untilTime))
}
