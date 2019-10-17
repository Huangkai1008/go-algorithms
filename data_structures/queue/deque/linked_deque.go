package deque

import (
	"fmt"
	"strings"
)

// DoubleNode 双向链表的节点
type DoubleNode struct {
	data interface{}
	prev *DoubleNode
	next *DoubleNode
}

// LinkedDeque 使用链表实现双端队列
//
// 使用链表头作为队首，链表尾作为队尾
// Front(Head): 2 <-> 4 <-> 1 ... <-> 3 Rear(Tail)
type LinkedDeque struct {
	front *DoubleNode
	rear  *DoubleNode
	size  int
}

// NewLinkedDeque 返回一个链表实现的双端队列
func NewLinkedDeque() *LinkedDeque {
	return &LinkedDeque{}
}

func (q *LinkedDeque) OfferFirst(data interface{}) {
	node := &DoubleNode{data: data}
	if q.IsEmpty() {
		q.front = node
		q.rear = node
	} else {
		q.front.prev = node
		node.next = q.front
		q.front = node
	}
	q.size++
}

func (q *LinkedDeque) OfferLast(data interface{}) {
	node := &DoubleNode{data: data}
	if q.IsEmpty() {
		q.front = node
		q.rear = node
	} else {
		q.rear.next = node
		node.prev = q.rear
		q.rear = node
	}
	q.size++
}

func (q *LinkedDeque) PopFirst() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty deque.")
	}

	data := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	} else {
		q.front.prev = nil
	}
	q.size--

	return data
}

func (q *LinkedDeque) PopLast() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty deque.")
	}

	data := q.rear.data
	q.rear = q.rear.prev
	if q.rear == nil {
		q.front = nil
	} else {
		q.rear.next = nil
	}
	q.size--

	return data
}

func (q *LinkedDeque) PeekFirst() interface{} {
	if q.IsEmpty() {
		panic("Can't peek from empty deque.")
	}

	return q.front.data
}

func (q *LinkedDeque) PeekLast() interface{} {
	if q.IsEmpty() {
		panic("Can't peek from empty deque.")
	}

	return q.rear.data
}

func (q *LinkedDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedDeque) Size() int {
	return q.size
}

func (q *LinkedDeque) String() string {
	var builder strings.Builder
	builder.WriteString("Front: ")
	for cur := q.front; cur != nil; cur = cur.next {
		builder.WriteString(fmt.Sprintf("%v<->", cur.data))
	}
	builder.WriteString(" Rear")
	return builder.String()
}
