package heap

import . "go-algorithms/data_structures/types"

type Option func(*MaxHeap)

// MaxHeap 最大堆的实现
type MaxHeap struct {
	items []T // 堆的数组存储
}

func NewMaxHeapWithItems(items []T) Option {
	return func(h *MaxHeap) {
		h.items = items
		h.heapify()
	}
}

// NewMaxHeap 返回一个新的最大堆
func NewMaxHeap(opts ...Option) *MaxHeap {
	heap := MaxHeap{
		items: []T{},
	}

	for _, o := range opts {
		o(&heap)
	}
	return &heap
}

func (h *MaxHeap) Push(item T) {
	h.items = append(h.items, item)
	h.siftUp(h.Size() - 1)
}

func (h *MaxHeap) Pop() T {
	top := h.Top()
	h.swap(0, h.Size()-1)
	h.items = h.items[:h.Size()-1]
	h.siftDown(0)
	return top
}

func (h *MaxHeap) Top() T {
	if h.IsEmpty() {
		panic("Can't get top from empty heap.")
	}

	return h.items[0]
}

func (h *MaxHeap) Replace(item T) T {
	top := h.Top()
	h.items[0] = item
	h.siftDown(0)
	return top
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.items) == 0
}

func (h *MaxHeap) Size() int {
	return len(h.items)
}

// parent 返回完全二叉树的数组中，给定元素所表示的元素父亲节点的索引
func (h *MaxHeap) parent(index int) int {
	if index <= 0 {
		panic("Get index failed, index out of range.")
	}

	return (index - 1) / 2
}

// left 返回完全二叉树的数组中，给定索引所表示的元素左孩子节点的索引
func (h *MaxHeap) left(index int) int {
	return index*2 + 1
}

// right 返回完全二叉树的数组中，给定索引所表示的元素右孩子节点的索引
func (h *MaxHeap) right(index int) int {
	return index*2 + 2
}

// heapify 将数组堆化
//
// 从最后一个叶子节点开始向前遍历下沉
// 最后一个叶子节点根据二叉堆的性质必定是最后一个节点的父节点
func (h *MaxHeap) heapify() {
	for i := h.parent(h.Size() - 1); i >= 0; i-- {
		h.siftDown(i)
	}
}

// swap 交换堆的两个元素
func (h *MaxHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// siftUp 堆的上浮操作
func (h *MaxHeap) siftUp(k int) {
	for ; k > 0 && h.items[h.parent(k)] < h.items[k]; k = h.parent(k) {
		h.swap(h.parent(k), k)
	}
}

// siftDown 堆的下沉操作
func (h *MaxHeap) siftDown(k int) {
	j := h.left(k)
	for length := h.Size(); j < length; j = h.left(k) {
		if j+1 < h.Size() && h.items[j+1] > h.items[j] {
			j = h.right(k)
		}

		if h.items[k] > h.items[j] {
			break
		}

		h.swap(j, k)
		k = j
	}
}
