package queue

import (
	"fmt"
	"strings"
)

// Node 链表的节点
type Node struct {
	data interface{}
	next *Node
}

// LinkedQueue 使用链表实现队列
//
// 使用链表头作为队首，链表尾作为队尾
// Front(Head): 2 -> 4 -> 1 ... -> 3 Rear(Tail)
type LinkedQueue struct {
	front *Node // 队列的头节点
	rear  *Node // 队列的尾结点
	size  int   // 队列的长度
}

// NewLinkedQueue 返回一个空的链表实现的队列
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func (q *LinkedQueue) Enqueue(data interface{}) {
	node := &Node{data: data}
	if q.IsEmpty() {
		q.front = node
		q.rear = node
	} else {
		q.rear.next = node
		q.rear = node
	}
	q.size++
}

func (q *LinkedQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Can't dequeue from empty queue.")
	}

	data := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
	return data
}

func (q *LinkedQueue) GetFront() interface{} {
	if q.IsEmpty() {
		panic("Can't get front from empty queue.")
	}
	return q.front.data
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedQueue) Size() int {
	return q.size
}

func (q *LinkedQueue) String() string {
	var builder strings.Builder
	builder.WriteString("Front: ")
	for cur := q.front; cur != nil; cur = cur.next {
		builder.WriteString(fmt.Sprintf("%v-->", cur.data))
	}
	builder.WriteString(" Rear")
	return builder.String()
}
