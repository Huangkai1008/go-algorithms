package arrayqueue

import arr "go-algorithms/data_structures/array"

// ArrayQueue 使用动态数组实现队列
//
// 使用数组头作为队列头，数组尾作为队列尾
// 示意图为 Front [1, 2, 3, ..., 9] Rear
type ArrayQueue struct {
	items *arr.DynamicArray
}

// New 返回一个空的动态数组实现的队列
func New() *ArrayQueue {
	return &ArrayQueue{arr.New()}
}

func (q *ArrayQueue) Enqueue(data interface{}) {
	q.items.Append(data)
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		panic("Can't dequeue from empty queue.")
	}

	return q.items.Remove(0)
}

func (q *ArrayQueue) GetFront() interface{} {
	if q.IsEmpty() {
		panic("Can't get front from empty queue.")
	}

	return q.items.Get(0)
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.items.IsEmpty()
}

func (q *ArrayQueue) Size() int {
	return q.items.Size()
}
