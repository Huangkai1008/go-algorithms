package linkedlist

import (
	"fmt"
	"strings"
)

// CircularLinkedList 环形链表的实现
type CircularLinkedList struct {
	*SingleLinkedList
}

// NewCll 返回一个空的环形链表
func NewCll() *CircularLinkedList {
	l := &CircularLinkedList{
		SingleLinkedList: NewSll(),
	}
	l.dummyHead.next = l.dummyHead
	return l
}

func (l *CircularLinkedList) Delete(data interface{}) {
	for pre := l.dummyHead; pre.next != l.dummyHead; pre = pre.next {
		cur := pre.next
		if cur.data == data {
			pre.next = cur.next
			cur.next = nil
			l.size--
		}
	}
}

func (l *CircularLinkedList) Contains(data interface{}) bool {
	for cur := l.dummyHead.next; cur != l.dummyHead; cur = cur.next {
		if cur.data == data {
			return true
		}
	}
	return false
}

func (l *CircularLinkedList) String() string {
	var builder strings.Builder
	builder.WriteString("Head: ")
	for cur := l.dummyHead.next; cur != l.dummyHead; cur = cur.next {
		builder.WriteString(fmt.Sprintf("%v-->", cur.data))
	}
	builder.WriteString("Head")
	return builder.String()
}
