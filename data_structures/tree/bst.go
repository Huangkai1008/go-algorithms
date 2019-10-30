package tree

type T int

type Node struct {
	data  T     // 节点的值
	left  *Node // 左子节点
	right *Node // 右子节点
}

// BST 二分搜索树的实现
//
// 二分搜索树是二叉树
//
// 二分搜索树的每个节点的值
// 	 * 大于其左子树的所有节点的值
//   * 小于其右子树的所有节点的值
//
// References: https://en.wikipedia.org/wiki/Binary_search_tree
type BST struct {
	root *Node // 二分搜索树的根
	size int   // 二分搜索树的节点个数
}

// NewBST 返回一个空的二分搜索树
func NewBST() *BST {
	return &BST{}
}

func (t *BST) IsEmpty() bool {
	return t.size == 0
}

func (t *BST) Size() int {
	return t.size
}
