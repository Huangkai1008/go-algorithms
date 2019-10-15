package queue

var (
	DefaultCapacity = 10
)

// Option 循环队列初始化的选项函数
type Option func(*CircularArrayQueue)

// CircularArrayQueue 使用动态数组实现循环队列
type CircularArrayQueue struct {
	front    int           // 队首的索引
	rear     int           // 队尾的索引
	size     int           // 队列的长度
	capacity int           // 队列的容量
	items    []interface{} // 队列的底层数组存储
}

// WithCapacity 设置队列的初始容量
func WithCapacity(capacity int) Option {
	return func(q *CircularArrayQueue) {
		q.capacity = capacity + 1
	}
}

// NewCircularArrayQueue 返回数组实现的空循环队列
func NewCircularArrayQueue(opts ...Option) *CircularArrayQueue {
	q := CircularArrayQueue{
		size:     0,
		capacity: DefaultCapacity + 1,
	}

	for _, o := range opts {
		o(&q)
	}
	q.items = make([]interface{}, q.capacity)
	return &q
}

func (q *CircularArrayQueue) Enqueue(data interface{}) {
	if q.isFull() {
		q.resize(q.getUseCapacity() * 2)
	}

	q.items[q.rear] = data
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

func (q *CircularArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Can't dequeue from empty queue.")
	}

	data := q.items[q.front]
	q.items[q.front] = nil
	q.front = (q.front + 1) % q.capacity
	q.size--

	if q.size == q.getUseCapacity()/4 && q.getUseCapacity()/2 != 0 {
		q.resize(q.getUseCapacity() / 2)
	}
	return data
}

func (q *CircularArrayQueue) GetFront() interface{} {
	if q.IsEmpty() {
		panic("Can't get front from empty queue.")
	}

	return q.items[q.front]
}

func (q *CircularArrayQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *CircularArrayQueue) Size() int {
	return q.size
}

// getUseCapacity 获取最大可用容量
func (q *CircularArrayQueue) getUseCapacity() int {
	return q.capacity - 1
}

// isFull 判断队列是否已满
func (q *CircularArrayQueue) isFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

// resize 队列扩缩容
func (q *CircularArrayQueue) resize(capacity int) {
	newItems := make([]interface{}, capacity+1)
	for i := 0; i < q.size; i++ {
		newItems[i] = q.items[(i+q.front)%q.capacity]
	}
	q.items = newItems
	q.front = 0
	q.rear = q.size
}
