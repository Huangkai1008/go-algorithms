// Package hashing 提供哈希表的定义与实现
//
// Ref: https://en.wikipedia.org/wiki/Hash_table
package hashing

import . "go-algorithms/data_structures/types"

const (
	DefaultCapacity        = 997  // 哈希表的默认容量
	DefaultUpperLoadFactor = 0.75 // 哈希表默认最大负载因子
	DefaultLowerLoadFactor = 0.10 // 哈希表默认最小负载因子
)

// HashTable 是哈希表的抽象类型
type HashTable interface {
	Put(key Key, value Value)   // Put 将键值对存入哈希表中
	Get(key Key) (Value, error) // Get 获取键对应的值，当键不存在时则返回keyError错误
	Delete(key Key)             // Delete 根据键删除对应的键值对
	Contains(key Key) bool      // Contains 判断键在表中是否有对应的值
	IsEmpty() bool              // IsEmpty 判断哈希表是否为空
	Size() int64                // Size 返回哈希表中元素的数量
}
