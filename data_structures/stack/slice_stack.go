package stack

import (
	"fmt"
	"strings"
)

// SliceStack 使用切片实现栈
//
// 使用切片末尾作为栈顶
// 示意图为 Bottom [..., 1, 3, 5 ] Top
type SliceStack struct {
	items []interface{}
}

// NewSliceStack 返回一个新的切片实现的空栈
func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func (s *SliceStack) Push(data interface{}) {
	s.items = append(s.items, data)
}

func (s *SliceStack) Pop() interface{} {
	if len(s.items) == 0 {
		panic("Can't pop from empty stack.")
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v
}

func (s *SliceStack) Peek() interface{} {
	if len(s.items) == 0 {
		panic("Can't peek from empty stack.")
	}

	return s.items[len(s.items)-1]
}

func (s *SliceStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *SliceStack) Size() int {
	return len(s.items)
}

func (s *SliceStack) String() string {
	var builder strings.Builder
	builder.WriteString("Bottom [")
	for i, item := range s.items {
		builder.WriteString(fmt.Sprint(item))
		if i != len(s.items)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("] Top")
	return builder.String()
}
