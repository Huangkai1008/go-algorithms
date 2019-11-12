// Package set 提供集合的实现
//
// Ref: References: https://en.wikipedia.org/wiki/Set_(abstract_data_type)
package set

import . "go-algorithms/data_structures/types"

// Set 是集合的抽象实现
type Set interface {
	Add(data T)           // Add 添加元素到集合中
	Remove(data T)        // Remove 从集合中删除元素
	Contains(data T) bool // Contains 判断集合中是否存在元素
	IsEmpty() bool        // IsEmpty 判断集合是否为空
	Size() int            // Size 返回集合中元素的个数
}
