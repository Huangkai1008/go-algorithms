// Package sorts 提供各种排序算法的实现
//
// 例如基础排序中的 SelectionSort、 InsertionSort
//
package sorts

type (
	Sort   func(items []int)
	Sorted func(items []int) []int
)
