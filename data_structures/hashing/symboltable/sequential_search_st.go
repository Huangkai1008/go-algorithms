package st

import "errors"

type Node struct {
	key   Key
	value Value
	next  *Node
}

// SequentialSearchST 顺序查找符号表
//
// 基于单向的无序链表实现
type SequentialSearchST struct {
	dummyHead *Node // 链表虚拟头节点
	size      int   // 符号表键值对的长度
}

// NewSequentialSearchST 返回新的顺序查找符号表
func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{
		dummyHead: &Node{},
	}
}

func (st *SequentialSearchST) Put(key Key, value Value) {
	for cur := st.dummyHead.next; cur != nil; cur = cur.next {
		if cur.key == key {
			cur.value = value
			return
		}
	}

	st.dummyHead.next = &Node{
		key:   key,
		value: value,
		next:  st.dummyHead.next,
	}
	st.size++
}

func (st *SequentialSearchST) Get(key Key) (Value, error) {
	for cur := st.dummyHead.next; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur.value, nil
		}
	}
	return 0, errors.New("keyError")
}

func (st *SequentialSearchST) Delete(key Key) {
	for pre, cur := st.dummyHead, st.dummyHead.next; cur != nil; pre, cur = pre.next, cur.next {
		if cur.key == key {
			pre.next = cur.next
			cur.next = nil
			st.size--
		}
	}
}

func (st *SequentialSearchST) Contains(key Key) bool {
	for cur := st.dummyHead.next; cur != nil; cur = cur.next {
		if cur.key == key {
			return true
		}
	}
	return false
}

func (st *SequentialSearchST) IsEmpty() bool {
	return st.size == 0
}

func (st *SequentialSearchST) Size() int {
	return st.size
}
