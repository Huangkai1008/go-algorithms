package stack

import (
	"fmt"
	"strings"

	ll "go-algorithms/data_structures/linkedlist"
)

// LinkedListStack 使用链表实现栈
//
// 使用链表头作为栈顶
// 示意图为 Top(Head) [ 1, 3, 5 ...] Bottom(Tail)
type LinkedListStack struct {
	items *ll.SingleLinkedList
}

// NewLinkedStack 返回一个新的单链表实现的空栈
func NewLinkedStack() *LinkedListStack {
	return &LinkedListStack{
		items: ll.NewSll(),
	}
}

func (s *LinkedListStack) Push(data interface{}) {
	s.items.Add(data)
}

func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		panic("Can't pop from empty stack.")
	}

	return s.items.Remove(0)
}

func (s *LinkedListStack) Peek() interface{} {
	if s.IsEmpty() {
		panic("Can't peek from empty stack.")
	}

	return s.items.Get(0)
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.items.IsEmpty()
}

func (s *LinkedListStack) Size() int {
	return s.items.Size()
}

func (s *LinkedListStack) String() string {
	var builder strings.Builder
	builder.WriteString("Top [ ")
	for i, item := range s.items.Values() {
		builder.WriteString(fmt.Sprint(item))
		if i != s.items.Size()-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(" ] Bottom")
	return builder.String()
}
