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

// Add 添加新的节点到二分搜索树中
func (t *BST) Add(data T) {
	t.root = t.add(t.root, data)
}

// add 插入元素到以node为根的二分搜索树中
// 返回插入新节点后二分搜索树的根
func (t *BST) add(node *Node, data T) *Node {
	if node == nil {
		t.size++
		return &Node{data: data}
	}

	if data < node.data {
		node.left = t.add(node.left, data)
	} else if data > node.data {
		node.right = t.add(node.right, data)
	}
	return node
}