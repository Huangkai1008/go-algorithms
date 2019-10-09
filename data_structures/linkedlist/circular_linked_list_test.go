package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularLinkedList(t *testing.T) {
	linkedList := NewCll()
	assert.Equal(t, true, linkedList.IsEmpty())
	fmt.Println(linkedList)

	for i := 0; i < 10; i++ {
		linkedList.Append(i)
	}
	fmt.Println(linkedList)

	assert.Equal(t, 10, linkedList.Size())
	assert.Equal(t, true, linkedList.Contains(0))
	assert.Equal(t, false, linkedList.Contains(11))
	assert.Equal(t, 0, linkedList.Remove(0))

	for i := 10; i < 20; i++ {
		linkedList.Append(i)
	}
	linkedList.Insert(2, 99)

	for i := 0; i < 10; i++ {
		linkedList.Remove(2)
	}
	linkedList.Remove(0)
	linkedList.Remove(2)
	linkedList.Delete(8)
	linkedList.Delete(18)
	linkedList.Delete(3)
	linkedList.Delete(1)

	for i := 0; i < linkedList.Size(); i++ {
		linkedList.Remove(0)
	}
	fmt.Println(linkedList)
}
