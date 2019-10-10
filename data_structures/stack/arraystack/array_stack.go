package arraystack

import (
	"strings"

	arr "go-algorithms/data_structures/array"
)

// ArrayStack 使用动态数组实现栈
//
// 使用数组末尾作为栈顶
// 示意图为 Bottom [..., 1, 3, 5 ] Top
type ArrayStack struct {
	items *arr.DynamicArray
}

// New 返回一个新的动态数组实现的空栈
func New() *ArrayStack {
	return &ArrayStack{
		items: arr.New(),
	}
}

func (s *ArrayStack) Push(data interface{}) {
	s.items.Append(data)
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		panic("Can't pop from empty stack.")
	}

	return s.items.Remove(s.items.Size() - 1)
}

func (s *ArrayStack) Peek() interface{} {
	if s.IsEmpty() {
		panic("Can't peek from empty stack.")
	}

	return s.items.Get(s.items.Size() - 1)
}

func (s *ArrayStack) IsEmpty() bool {
	return s.items.IsEmpty()
}

func (s *ArrayStack) Size() int {
	return s.items.Size()
}

func (s *ArrayStack) String() string {
	var builder strings.Builder
	builder.WriteString("Bottom ")
	builder.WriteString(s.items.String())
	builder.WriteString(" Top")
	return builder.String()
}
