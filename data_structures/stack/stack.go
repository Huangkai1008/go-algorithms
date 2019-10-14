// Package stack 提供队列的定义与实现
// 栈是一种后进先出（LIFO）的数据结构
//
// Ref: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
package stack

// Stack 是栈的抽象类型
type Stack interface {
	Push(data interface{}) // Push 添加元素到栈顶
	Pop() interface{}      // Pop 弹出栈顶元素
	Peek() interface{}     // Peek 查看栈顶元素
	IsEmpty() bool         // IsEmpty 判断栈是否为空
	Size() int             // Size 返回栈内元素的数量
}
