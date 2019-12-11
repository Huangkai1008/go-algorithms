package heap

import . "go-algorithms/data_structures/types"

// Heap 堆的抽象实现
type Heap interface {
	Push(item T)      // Push 添加元素到堆中
	Pop() T           // Pop 弹出堆顶元素
	Top() T           // Top 返回堆顶元素
	Replace(item T) T // Replace 替换并返回堆顶元素
	IsEmpty() bool    // IsEmpty 判断堆是否为空
	Size() int        // Size 返回堆中元素的数量
}
