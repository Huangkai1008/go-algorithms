package singlelinkedlist

// Node 链表的节点
type Node struct {
	data interface{}
	next *Node
}

// SingleLinkedList 单链表实现
type SingleLinkedList struct {
	dummyHead *Node // 链表的虚拟头节点
	size      int   // 链表的长度
}

func New() *SingleLinkedList {
	return &SingleLinkedList{
		dummyHead: &Node{},
		size:      0,
	}
}

func (l *SingleLinkedList) Add(data interface{}) {
	l.Set(0, data)
}

func (l *SingleLinkedList) Append(data interface{}) {
	l.Set(l.size, data)
}

func (l *SingleLinkedList) Set(pos int, data interface{}) {
	pre := l.dummyHead
	for i := 0; i < pos; i++ {
		pre = pre.next
	}

	node := &Node{data: data}
	node.next = pre.next
	pre.next = node
	l.size++
}

func (l *SingleLinkedList) Get(pos int) interface{} {
	if pos >= l.size {
		panic("Get failed, position out of range.")
	}

	cur := l.dummyHead
	for i := 0; i < pos+1; i++ {
		cur = cur.next
	}
	return cur
}

func (l *SingleLinkedList) Remove(pos int) interface{} {
	if pos >= l.size {
		panic("Remove failed, position out of range.")
	}

	pre := l.dummyHead
	for i := 0; i < pos; i++ {
		pre = pre.next
	}

	cur := pre.next
	pre.next = cur.next
	cur.next = nil
	l.size--
	return cur.data
}

func (l *SingleLinkedList) Delete(data interface{}) {
	pre := l.dummyHead
	for pre.next != nil {
		cur := pre.next
		if cur.data == data {
			pre.next = cur.next
			cur.next = nil
			l.size--
		}

		pre = pre.next
	}
}

func (l *SingleLinkedList) Contains(data interface{}) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if cur.data == data {
			return true
		}
		cur = cur.next
	}

	return false
}

func (l *SingleLinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *SingleLinkedList) Size() int {
	return l.size
}
