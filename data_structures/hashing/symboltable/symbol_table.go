// Package st 提供符号表的定义与实现
//
// 符号表是一种存储键值对的数据结构
//
//
// Ref: https://en.wikipedia.org/wiki/Symbol_table
package st

type (
	Key   string // 符号表的键类型
	Value string // 符号表的值类型
)

// SymbolTable 是符号表的抽象类型
type SymbolTable interface {
	Put(key Key, value Value)  // Put 将键值对存入符号表中
	Get(key Key) (Value error) // Get 获取键对应的值，当键不存在时则返回keyError错误
	Delete(key Key)            // Delete 根据键删除对应的键值对
	Contains(key Key) bool     // Contains 判断键在表中是否有对应的值
	IsEmpty() bool             // IsEmpty 判断符号表是否为空
	Size() int                 // Size 返回符号表中元素的数量
}
