package deque

import (
	"fmt"
	"strings"
)

var (
	DefaultCapacity = 10
)

// Option 循环双端队列初始化的选项
type Option func(*CircularDeque)

// CircularDeque 数组实现循环双端队列
type CircularDeque struct {
	front    int           // 队首的索引
	rear     int           // 队尾的索引
	size     int           // 队列的长度
	capacity int           // 队列的容量
	items    []interface{} // 队列的底层数组
}

// WithCapacity 设置双端队列的初始容量
func WithCapacity(capacity int) Option {
	return func(q *CircularDeque) {
		q.capacity = capacity + 1
	}
}

func NewCircularDeque(opts ...Option) *CircularDeque {
	q := CircularDeque{
		size:     0,
		capacity: DefaultCapacity + 1,
	}
	for _, o := range opts {
		o(&q)
	}
	q.items = make([]interface{}, q.capacity)
	return &q
}

func (q *CircularDeque) OfferFirst(data interface{}) {
	if q.isFull() {
		q.resize(q.capacity * 2)
	}

	q.front = (q.front - 1 + q.capacity) % q.capacity
	q.items[q.front] = data
	q.size++
}

func (q *CircularDeque) OfferLast(data interface{}) {
	if q.isFull() {
		q.resize(q.capacity * 2)
	}

	q.items[q.rear] = data
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

func (q *CircularDeque) PopFirst() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty deque.")
	}

	data := q.items[q.front]
	q.items[q.front] = nil
	q.front = (q.front + 1) % q.capacity
	q.size--

	if q.size == q.useCapacity()/4 && q.useCapacity()/2 != 0 {
		q.resize(q.useCapacity() / 2)
	}
	return data
}

func (q *CircularDeque) PopLast() interface{} {
	if q.IsEmpty() {
		panic("Can't pop from empty deque.")
	}

	data := q.items[(q.rear-1+q.capacity)%q.capacity]
	q.items[(q.rear-1+q.capacity)%q.capacity] = nil
	q.rear = (q.rear - 1 + q.capacity) % q.capacity
	q.size--

	if q.size == q.useCapacity()/4 && q.useCapacity()/2 != 0 {
		q.resize(q.useCapacity() / 2)
	}
	return data
}

func (q *CircularDeque) PeekFirst() interface{} {
	if q.IsEmpty() {
		panic("Can't peek from empty deque.")
	}

	return q.items[q.front]
}

func (q *CircularDeque) PeekLast() interface{} {
	if q.IsEmpty() {
		panic("Can't peek from empty deque.")
	}

	return q.items[(q.rear-1)%q.capacity]
}

func (q *CircularDeque) IsEmpty() bool {
	return q.front == q.rear
}

func (q *CircularDeque) Size() int {
	return q.size
}

// useCapacity 返回最大可用容量
func (q *CircularDeque) useCapacity() int {
	return len(q.items) - 1
}

// isFull 判断队列是否为空
func (q *CircularDeque) isFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

// resize 队列扩缩容
func (q *CircularDeque) resize(capacity int) {
	newItems := make([]interface{}, capacity)
	for i := 0; i < q.size; i++ {
		newItems[i] = q.items[(i+q.front)%q.capacity]
	}
	q.items = newItems
	q.front = 0
	q.rear = q.size
	q.capacity = capacity + 1
}

func (q *CircularDeque) String() string {
	var builder strings.Builder
	builder.WriteString("Front: [ ")
	for i := q.front; i != q.rear; i = (i + 1) % q.capacity {
		builder.WriteString(fmt.Sprintf("%v ", q.items[i]))
	}
	builder.WriteString("] Rear")
	return builder.String()
}
