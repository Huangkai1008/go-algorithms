// Package tree 提供树的定义与实现
package tree

type Tree interface {
	IsEmpty() bool // IsEmpty 判断树是否为空
	Size() int     // Size 返回树中节点的个数
}
