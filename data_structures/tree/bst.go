package tree

import "fmt"

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

// Contains 判断元素是否在二分搜索树中
func (t *BST) Contains(data T) bool {
	return t.contains(t.root, data)
}

// PreOrder 二分搜索树的前序遍历
//
// 遍历顺序：根节点 --> 左子树 --> 右子树
func (t *BST) PreOrder() {
	t.preOrder(t.root)
}

// InOrder 二分搜索树的中序遍历
//
// 遍历顺序：左子树 --> 根节点 --> 右子树
func (t *BST) InOrder() {
	t.inOrder(t.root)
}

// PostOrder 二分搜索树的后序遍历
//
// 遍历顺序：左子树 --> 右子树 --> 根节点
func (t *BST) PostOrder() {
	t.postOrder(t.root)
}

// PreOrderNR 前序遍历非递归实现
func (t *BST) PreOrderNR() {
	if t.IsEmpty() {
		return
	}

	var stack []*Node
	stack = append(stack, t.root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		fmt.Print(cur.data, " ")

		stack = stack[:len(stack)-1] // 出栈

		// 先压入右子树再压入左子树
		if cur.right != nil {
			stack = append(stack, cur.right)
		}
		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}
}

// LevelOrder 层序遍历非递归实现
func (t *BST) LevelOrder() {
	if t.IsEmpty() {
		return
	}

	var queue []*Node
	queue = append(queue, t.root)
	for len(queue) > 0 {
		cur := queue[0]
		fmt.Print(cur.data, " ")

		queue = queue[1:]

		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
}

// Min 返回最小的元素
func (t *BST) Min() T {
	if t.IsEmpty() {
		panic("Get min node from empty bst.")
	}

	return t.min(t.root).data
}

// Max 返回最大的元素
func (t *BST) Max() T {
	if t.IsEmpty() {
		panic("Get max node from empty bst.")
	}

	return t.max(t.root).data
}

// RemoveMin 从二分搜索树中删除最小值所在节点
// 返回最小值
func (t *BST) RemoveMin() T {
	min := t.Min()
	t.removeMin(t.root)
	return min
}

// RemoveMax 从二分搜索树中删除最大值所在节点
// 返回最大值
func (t *BST) RemoveMax() T {
	max := t.Max()
	t.removeMax(t.root)
	return max
}

// Remove 从二分搜索树中移除给定的节点
func (t *BST) Remove(data T) {
	t.root = t.remove(t.root, data)
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

// contains 判断元素是否在以node为根的二分搜索树中
func (t *BST) contains(node *Node, data T) bool {
	if node == nil {
		return false
	}

	if data == node.data {
		return true
	} else if data < node.data {
		return t.contains(node.left, data)
	} else {
		return t.contains(node.right, data)
	}
}

// preOrder 前序遍历以Node为根的二分搜索树
func (t *BST) preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.data, " ")
	t.preOrder(node.left)
	t.preOrder(node.right)
}

// inOrder 中序遍历以Node为根的二分搜索树
func (t *BST) inOrder(node *Node) {
	if node == nil {
		return
	}

	t.inOrder(node.left)
	fmt.Print(node.data, " ")
	t.inOrder(node.right)
}

// postOrder 中序遍历以Node为根的二分搜索树
func (t BST) postOrder(node *Node) {
	if node == nil {
		return
	}

	t.postOrder(node.left)
	t.postOrder(node.right)
	fmt.Print(node.data, " ")
}

// min 返回以node为根的二分搜索树的最小元素所在的节点
func (t *BST) min(node *Node) *Node {
	if node.left == nil {
		return node
	}

	return t.min(node.left)
}

// max 返回以node为根的二分搜索树的最大元素所在的节点
func (t *BST) max(node *Node) *Node {
	if node.right == nil {
		return node
	}

	return t.max(node.right)
}

// removeMin 从以node为根的二分搜索树中删除最小节点
// 返回删除节点后新的二分搜索树的根节点
func (t *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		right := node.right
		node.right = nil
		t.size--
		return right
	}

	node.left = t.removeMin(node.left)
	return node
}

// removeMax 从以node为根的二分搜索树中删除最大节点
// 返回删除节点后新的二分搜索树的根节点
func (t *BST) removeMax(node *Node) *Node {
	if node.right == nil {
		left := node.left
		node.left = nil
		t.size--
		return left
	}

	node.right = t.removeMax(node.right)
	return node
}

// remove 从以node为根的二分搜索树中删除给定节点
// 返回删除节点后新的二分搜索树的根节点
func (t *BST) remove(node *Node, data T) *Node {
	if node == nil {
		return nil
	}

	if data < node.data {
		node.left = t.remove(node.left, data)
		return node
	} else if data > node.data {
		node.right = t.remove(node.right, data)
		return node
	} else {
		if node.left == nil {
			// 待删除的节点左子树为空
			right := node.right
			node.right = nil
			t.size--
			return right
		}
		if node.right == nil {
			// 待删除的节点右子树为空
			left := node.left
			node.left = nil
			t.size--
			return left
		}

		// 待删除节点左子树和右子树均不为空的情况
		// 找到比这个待删除节点大的最小节点
		// 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := t.min(node.right)
		successor.right = t.removeMin(node.right)
		successor.left = node.left
		node.left = nil
		node.right = nil
		return successor
	}
}
