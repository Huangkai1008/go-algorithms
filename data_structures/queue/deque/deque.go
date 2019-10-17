// Package deque 提供双端队列的定义和实现
//
// Ref: https://en.wikipedia.org/wiki/Double-ended_queue
package deque

// Deque 是双端队列的抽象类型
type Deque interface {
	OfferFirst(data interface{}) // OfferFirst 添加元素到队首
	OfferLast(data interface{})  // OfferLast 添加元素到队尾
	PopFirst() interface{}       // PopFirst 队首元素出队
	PopLast() interface{}        // PopLast 队尾元素出队
	PeekFirst() interface{}      // PeekFirst 查看队首元素
	PeekLast() interface{}       // PeekLast 查看队尾元素
	IsEmpty() bool               // IsEmpty 判断队列是否为空
	Size() int                   // Size 返回队列的长度
}
