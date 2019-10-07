package slicestack

// SliceStack 使用切片实现栈
//
// 使用切片末尾作为栈顶
// 示意图为 Bottom [..., 1, 3, 5 ] Top
type SliceStack struct {
	items []interface{}
}

// New 返回一个新的切片实现的空栈
func New() *SliceStack {
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
		return nil
	}

	return s.items[len(s.items)-1]
}

func (s SliceStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s SliceStack) Size() int {
	return len(s.items)
}
