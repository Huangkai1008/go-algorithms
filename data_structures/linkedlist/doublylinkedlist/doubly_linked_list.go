package dll

import (
	"fmt"
	"strings"
)

// DoublyLinkedNode 双向链表的节点
type DoublyLinkedNode struct {
	data interface{}
	prev *DoublyLinkedNode
	next *DoublyLinkedNode
}

// DoublyLinkedList 双向链表的实现
type DoublyLinkedList struct {
	head *DoublyLinkedNode // 双向链表的头部节点
	tail *DoublyLinkedNode // 双向链表的尾部节点
	size int               // 链表的长度
}

// NewDll 返回一个空的双向链表
func NewDll() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) Add(data interface{}) {
	l.Insert(0, data)
}

func (l *DoublyLinkedList) Append(data interface{}) {
	l.Insert(l.size, data)
}

func (l *DoublyLinkedList) Insert(pos int, data interface{}) {
	if pos < 0 || pos > l.size {
		panic("Insert failed, position out of range.")
	}

	node := &DoublyLinkedNode{data: data}
	if l.size == 0 {
		l.head = node
		l.tail = node
	} else if pos == 0 {
		node.next = l.head
		l.head.prev = node
		l.head = node
	} else if pos == l.size {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	} else {
		pre := l.head
		for i := 0; i < pos-1; i++ {
			pre = pre.next
		}
		node.next = pre.next
		pre.next.prev = node
		pre.next = node
		node.prev = pre
	}

	l.size++
}

func (l *DoublyLinkedList) Get(pos int) interface{} {
	if pos < 0 || pos >= l.size {
		panic("Get failed, position out of range.")
	}

	cur := l.head
	for i := 0; i < pos; i, cur = i+1, cur.next {
	}
	return cur.data
}

func (l *DoublyLinkedList) Remove(pos int) interface{} {
	if pos < 0 || pos >= l.size {
		panic("Remove failed, position out of range.")
	}

	if l.size == 1 {
		data := l.head.data
		l.head = nil
		l.tail = nil
		l.size = 0
		return data
	}

	cur := l.head
	for i := 0; i < pos; i++ {
		cur = cur.next
	}
	if cur == l.head {
		l.head = l.head.next
	}
	if cur == l.tail {
		l.tail = l.tail.prev
	}

	if cur.prev != nil {
		cur.prev.next = cur.next
	}
	if cur.next != nil {
		cur.next.prev = cur.prev
	}

	data := cur.data
	cur = nil
	l.size--
	return data
}

func (l *DoublyLinkedList) Delete(data interface{}) {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.data == data {
			if cur == l.head {
				l.head = l.head.next
				cur.next.prev = nil
				cur.next = nil
			} else if cur == l.tail {
				l.tail = l.tail.prev
				cur.prev.next = nil
				cur.prev = nil
			} else {
				cur.prev.next = cur.next
				cur.next.prev = cur.prev
				cur.prev = nil
				cur.next = nil
			}
			l.size--
			return
		}
	}
}

func (l *DoublyLinkedList) Contains(data interface{}) bool {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.data == data {
			return true
		}
	}
	return false
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *DoublyLinkedList) Size() int {
	return l.size
}

func (l *DoublyLinkedList) String() string {
	var builder strings.Builder
	builder.WriteString("Head: ")
	for cur := l.head; cur != nil; cur = cur.next {
		builder.WriteString(fmt.Sprintf("%v<--->", cur.data))
	}
	builder.WriteString("Null")
	return builder.String()
}
