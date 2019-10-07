package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	array := New(10)
	assert.Equal(t, true, array.IsEmpty())

	for i := 0; i < 19; i++ {
		array.Append(i)
	}
	assert.Equal(t, 20, array.Capacity())
	assert.Equal(t, 19, array.Size())

	array.Set(0, 1)
	assert.Equal(t, 1, array.Get(0))

	array.Insert(18, 99)
	assert.Equal(t, 18, array.Search(99))
	assert.Equal(t, -1, array.Search(100))
	assert.Equal(t, true, array.Contains(99))
	assert.Equal(t, false, array.Contains(100))

	array.Remove(0)
	array.Remove(1)
	fmt.Println(array)
}
