package array

import (
	"fmt"
	"strings"
)

const (
	DefaultCapacity = 10
)

// Option 数组初始化的选项函数
type Option func(*DynamicArray)

// DynamicArray 实现动态数组
type DynamicArray struct {
	capacity int
	size     int
	items    []interface{}
}

// WithCapacity 设置数组的初始容量
func WithCapacity(capacity int) Option {
	return func(a *DynamicArray) {
		a.capacity = capacity
	}
}

// New 返回一个空动态数组
func New(opts ...Option) *DynamicArray {
	arr := DynamicArray{
		capacity: DefaultCapacity,
		size:     0,
	}

	for _, o := range opts {
		o(&arr)
	}
	arr.items = make([]interface{}, arr.capacity)
	return &arr
}

func (a *DynamicArray) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get item failed, index out of range.")
	}

	return a.items[index]
}

func (a *DynamicArray) Append(item interface{}) {
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}

	a.items[a.size] = item
	a.size++
}

func (a *DynamicArray) Set(index int, item interface{}) {
	if index < 0 || index >= a.size {
		panic("Set item failed, index out of range.")
	}

	a.items[index] = item
}

func (a *DynamicArray) Insert(index int, item interface{}) {
	if index < 0 || index > a.size {
		panic("Insert item failed, index out of range.")
	}

	// 当数组的元素数量等于数组容量时需要进行扩容
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}

	for i := a.size - 1; i >= index; i-- {
		a.items[i+1] = a.items[i]
	}
	a.items[index] = item
	a.size++
}

func (a *DynamicArray) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Remove item failed, index out of range.")
	}

	v := a.items[index]
	for i := index; i < a.size-1; i++ {
		a.items[i] = a.items[i+1]
	}
	a.items[a.size-1] = nil
	a.size--

	// 当数组的元素数量等于数组容量的1/4时需要进行缩容
	// 同时需要考虑边界情况，缩容后的容量不能为0
	if a.size == a.capacity/4 && a.capacity/2 != 0 {
		a.resize(a.capacity / 2)
	}

	return v
}

func (a *DynamicArray) Contains(item interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.items[i] == item {
			return true
		}
	}
	return false
}

func (a *DynamicArray) Search(item interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.items[i] == item {
			return i
		}
	}
	return -1
}

func (a *DynamicArray) IsEmpty() bool {
	return a.size == 0
}

func (a *DynamicArray) Size() int {
	return a.size
}

func (a *DynamicArray) Capacity() int {
	return a.capacity
}

// resize 进行数组的扩缩容
func (a *DynamicArray) resize(capacity int) {
	newItems := make([]interface{}, capacity)
	copy(newItems, a.items)
	a.capacity = capacity
	a.items = newItems
}

func (a *DynamicArray) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for i := 0; i < a.size; i++ {
		builder.WriteString(fmt.Sprint(a.items[i]))
		if i != a.size-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}
