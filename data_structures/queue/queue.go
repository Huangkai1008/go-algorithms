// Package queue 提供队列的定义和实现
//
// 队列是一种先进先出（FIFO）的数据结构
//
// Ref: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
//
// 实现的单向队列有 ArrayQueue、 LinkedQueue
package queue

// Queue 是先进先出队列的抽象类型
type Queue interface {
	Enqueue(data interface{}) // Enqueue 入队
	Dequeue() interface{}     // Dequeue 出队
	GetFront() interface{}    // GetFront 获取队首元素
	IsEmpty() bool            // IsEmpty 判断队列是否为空
	Size() int                // Size 返回队列的长度
}
