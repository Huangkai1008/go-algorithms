package queue

import (
	"fmt"
	"strings"
)

const (
	DefaultCapacity = 10
)

// Option 循环队列初始化的选项函数
type Option func(*CircularQueue)

// CircularQueue 使用动态数组实现循环队列
type CircularQueue struct {
	front    int           // 队首的索引
	rear     int           // 队尾的索引
	size     int           // 队列的长度
	capacity int           // 队列的容量
	items    []interface{} // 队列的底层数组
}

// WithCapacity 设置队列的初始容量
func WithCapacity(capacity int) Option {
	return func(q *CircularQueue) {
		q.capacity = capacity + 1
	}
}

// NewCircularArrayQueue 返回数组实现的空循环队列
func NewCircularArrayQueue(opts ...Option) *CircularQueue {
	q := CircularQueue{
		size:     0,
		capacity: DefaultCapacity + 1,
	}

	for _, o := range opts {
		o(&q)
	}
	q.items = make([]interface{}, q.capacity)
	return &q
}

func (q *CircularQueue) Enqueue(data interface{}) {
	if q.isFull() {
		q.resize(q.useCapacity() * 2)
	}

	q.items[q.rear] = data
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Can't dequeue from empty queue.")
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

func (q *CircularQueue) GetFront() interface{} {
	if q.IsEmpty() {
		panic("Can't get front from empty queue.")
	}

	return q.items[q.front]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *CircularQueue) Size() int {
	return q.size
}

// useCapacity 返回最大可用容量
func (q *CircularQueue) useCapacity() int {
	return len(q.items) - 1
}

// isFull 判断队列是否已满
func (q *CircularQueue) isFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

// resize 队列扩缩容
func (q *CircularQueue) resize(capacity int) {
	newItems := make([]interface{}, capacity+1)
	for i := 0; i < q.size; i++ {
		newItems[i] = q.items[(i+q.front)%q.capacity]
	}
	q.items = newItems
	q.front = 0
	q.rear = q.size
	q.capacity = capacity + 1
}

func (q *CircularQueue) String() string {
	var builder strings.Builder
	builder.WriteString("Front: [")
	for i := q.front; i != q.rear; i = (i + 1) % q.capacity {
		builder.WriteString(fmt.Sprint(q.items[i]))
		if (i+1)%q.capacity != q.rear {
			builder.WriteString(" ,")
		}
	}
	builder.WriteString("] Rear")
	return builder.String()
}
