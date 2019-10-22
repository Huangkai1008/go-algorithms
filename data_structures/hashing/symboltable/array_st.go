package st

import "errors"

var DefaultCapacity = 10

type Option func(*ArrayST)

// ArrayST 平行数组实现符号表
type ArrayST struct {
	keys     []Key
	values   []Value
	size     int
	capacity int
}

// WithCapacity 设置符号表的初始容量
func WithCapacity(capacity int) Option {
	return func(st *ArrayST) {
		st.capacity = capacity
	}
}

// NewArrayST 返回一个空符号表
func NewArrayST(opts ...Option) *ArrayST {
	st := ArrayST{
		size:     0,
		capacity: DefaultCapacity,
	}

	for _, o := range opts {
		o(&st)
	}
	return &st
}

func (st *ArrayST) Put(key Key, value Value) {
	for i := 0; i < st.size; i++ {
		if st.keys[i] == key {
			st.values[i] = value
			return
		}
	}

	if st.size == st.capacity {
		st.resize(st.capacity * 2)
	}

	st.keys[st.size] = key
	st.values[st.size] = value
	st.size++
}

func (st *ArrayST) Get(key Key) (Value, error) {
	for i := 0; i < st.size; i++ {
		if st.keys[i] == key {
			return st.values[i], nil
		}
	}
	return "", errors.New("keyError")
}

func (st *ArrayST) Delete(key Key) {
	for i := 0; i < st.size; i++ {
		if st.keys[i] == key {
			st.keys = append(st.keys[:i], st.keys[i+1:]...)
			st.values = append(st.values[:i], st.values[i+1:]...)
			st.size--

			if st.size == st.capacity/4 && st.capacity/2 != 0 {
				st.resize(st.capacity >> 1)
			}
		}
	}
}

func (st *ArrayST) Contains(key Key) bool {
	for i := 0; i < st.size; i++ {
		if st.keys[i] == key {
			return true
		}
	}
	return false
}

func (st *ArrayST) IsEmpty() bool {
	return st.size == 0
}

func (st *ArrayST) Size() int {
	return st.size
}

// resize 符号表扩缩容
func (st *ArrayST) resize(capacity int) {
	newKeys, newValues := make([]Key, capacity), make([]Value, capacity)
	copy(newKeys, st.keys)
	copy(newValues, st.values)
	st.keys = newKeys
	st.values = newValues
	st.capacity = capacity
}
