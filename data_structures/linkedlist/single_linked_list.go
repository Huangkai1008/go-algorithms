package linkedlist

import (
	"fmt"
	"strings"
)

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

// NewSll 返回一个空的单链表
func NewSll() *SingleLinkedList {
	return &SingleLinkedList{
		dummyHead: &Node{},
	}
}

func (l *SingleLinkedList) Add(data interface{}) {
	l.Insert(0, data)
}

func (l *SingleLinkedList) Append(data interface{}) {
	l.Insert(l.size, data)
}

func (l *SingleLinkedList) Insert(pos int, data interface{}) {
	if pos < 0 || pos > l.size {
		panic("Insert failed, position out of range.")
	}

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
	if pos < 0 || pos >= l.size {
		panic("Get failed, position out of range.")
	}

	cur := l.dummyHead.next
	for i := 0; i < pos; i, cur = i+1, cur.next {
	}
	return cur.data
}

func (l *SingleLinkedList) Remove(pos int) interface{} {
	if pos < 0 || pos >= l.size {
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
	for pre := l.dummyHead; pre.next != nil; pre = pre.next {
		cur := pre.next
		if cur.data == data {
			pre.next = cur.next
			cur.next = nil
			l.size--
		}
	}
}

func (l *SingleLinkedList) Contains(data interface{}) bool {
	for cur := l.dummyHead.next; cur != nil; cur = cur.next {
		if cur.data == data {
			return true
		}
	}

	return false
}

func (l *SingleLinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *SingleLinkedList) Size() int {
	return l.size
}

// Values 返回所有节点的值列表
func (l *SingleLinkedList) Values() []interface{} {
	values := make([]interface{}, l.size)
	for i, cur := 0, l.dummyHead.next; cur != nil; i, cur = i+1, cur.next {
		values[i] = cur.data
	}
	return values
}

func (l *SingleLinkedList) String() string {
	var builder strings.Builder
	builder.WriteString("Head: ")
	for cur := l.dummyHead.next; cur != nil; cur = cur.next {
		builder.WriteString(fmt.Sprintf("%v-->", cur.data))
	}
	builder.WriteString("Null")
	return builder.String()
}
