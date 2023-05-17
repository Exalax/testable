package uuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewString(t *testing.T) {
	Reset()

	id := NewString()
	assert.NotEqual(t, uuid.UUID{}.String(), id)
}

func TestSet(t *testing.T) {
	id1 := uuid.New()
	id2 := uuid.New()

	Set(id1, id2)
	defer Reset()

	assert.Equal(t, id1, New())
	assert.Equal(t, id2, New())

	assert.NotEqual(t, uuid.UUID{}, New())
}

func TestSetRandom(t *testing.T) {
	ids := SetRandom(3)
	defer Reset()

	assert.Equal(t, ids, []uuid.UUID{New(), New(), New()})

	assert.NotEqual(t, uuid.UUID{}, New())
}

func TestReset(t *testing.T) {
	// Reset after Set.
	id := uuid.New()
	Set(id)
	Reset()

	assert.NotEqual(t, id, New())

	// Reset after SetRandom.
	ids := SetRandom(3)
	Reset()

	assert.NotEqual(t, ids[0], New())
}
