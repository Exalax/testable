package uuid

import "github.com/google/uuid"

var buf []uuid.UUID

func New() uuid.UUID {
	if len(buf) == 0 {
		return uuid.New()
	}

	id := buf[0]
	buf = buf[1:]

	return id
}

func NewString() string {
	return New().String()
}

func Reset() {
	buf = nil
}

func Set(ids ...uuid.UUID) {
	buf = ids
}

func SetRandom(n int) []uuid.UUID {
	buf = make([]uuid.UUID, n)
	uuids := make([]uuid.UUID, n)
	for i := range buf {
		buf[i] = uuid.New()
		uuids[i] = buf[i]
	}

	return uuids
}

func Parse(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
