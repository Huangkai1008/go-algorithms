package deque

import (
	arr "go-algorithms/data_structures/array"
	"strings"
)

// ArrayDeque 使用数组实现双端队列
//
// 使用数组头作为队首，数组尾作为队尾
// 示意图为 Front [1, 2, 3, ..., 9] Rear
type ArrayDeque struct {
	items *arr.DynamicArray
}

// NewArrayDeque 返回一个数组实现的空的双端队列
func NewArrayDeque() *ArrayDeque {
	return &ArrayDeque{
		items: arr.New(),
	}
}

func (q *ArrayDeque) OfferFirst(data interface{}) {
	q.items.Insert(0, data)
}

func (q *ArrayDeque) OfferLast(data interface{}) {
	q.items.Append(data)
}

func (q *ArrayDeque) PopFirst() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty deque.")
	}

	return q.items.Remove(0)
}

func (q *ArrayDeque) PopLast() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty deque.")
	}

	return q.items.Remove(q.Size() - 1)
}

func (q *ArrayDeque) PeekFirst() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty queue.")
	}

	return q.items.Get(0)
}

func (q *ArrayDeque) PeekLast() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty queue.")
	}

	return q.items.Get(q.Size() - 1)
}

func (q *ArrayDeque) IsEmpty() bool {
	return q.items.IsEmpty()
}

func (q *ArrayDeque) Size() int {
	return q.items.Size()
}

func (q *ArrayDeque) String() string {
	var builder strings.Builder
	builder.WriteString("Front: ")
	builder.WriteString(q.items.String())
	builder.WriteString(" Rear")
	return builder.String()
}
