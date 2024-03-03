package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMemoryEvent(t *testing.T) {
	e := MemoryEvent{}

	domainModel := toMemoryEvent(&e)

	assert.Equal(t, e.Id, domainModel.Id)
}

func TestToMemoryEvents(t *testing.T) {
	e := MemoryEvent{}

	es := []*MemoryEvent{&e}

	domainModels := toMemoryEvents(es)

	assert.Equal(t, len(domainModels), 1)
}
