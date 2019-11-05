// Package hashing 提供哈希表的定义与实现
//
// Ref: https://en.wikipedia.org/wiki/Hash_table
package hashing

import st "go-algorithms/data_structures/hashing/symboltable"

const (
	DefaultCapacity        = 997  // 哈希表的默认容量
	DefaultUpperLoadFactor = 0.75 // 哈希表默认最大负载因子
	DefaultLowerLoadFactor = 0.10 // 哈希表默认最小负载因子
)

// HashTable 是哈希表的抽象类型
type HashTable interface {
	st.SymbolTable
}
