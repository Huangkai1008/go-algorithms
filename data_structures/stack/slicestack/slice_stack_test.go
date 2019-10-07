package slicestack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceStack(t *testing.T) {
	stack := New()
	assert.Equal(t, true, stack.IsEmpty())
	fmt.Println(stack)
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	fmt.Println(stack)

	assert.Equal(t, 10, stack.Size())
	assert.Equal(t, 9, stack.Pop())
	assert.Equal(t, 8, stack.Peek())
	fmt.Println(stack)
}
