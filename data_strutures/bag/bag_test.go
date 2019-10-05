package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBag(t *testing.T) {
	bag := New()
	assert.Equal(t, true, bag.IsEmpty(), "The bag should be empty.")
	assert.Equal(t, 0, bag.Size(), "The empty bag's size should be zero.")

	bag.Add(1)
	bag.Add(2)
	bag.Add(3)
	assert.Equal(t, 3, bag.Size(), "The bag's size should be three now.")
	assert.Equal(t, false, bag.IsEmpty(), "The bag should be not empty.")
}
