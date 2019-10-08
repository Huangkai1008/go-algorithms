package linkedlist

// LinkedList 是链表的抽象类型
type LinkedList interface {
	Add(data interface{})             // Add 在链表头部添加新的节点
	Append(data interface{})          // Append 在链表尾部添加新的节点
	Insert(pos int, data interface{}) // Insert 设置链表给定位置节点的值
	Get(pos int) interface{}          // Get 返回链表给定位置节点的值
	Remove(pos int) interface{}       // Remove 移除链表给定位置的节点并返回节点的值
	Delete(data interface{})          // Delete 删除链表中符合给定值的节点
	Contains(data interface{}) bool   // Contains 判断节点是否在链表中

	IsEmpty() bool // IsEmpty 判断链表是否为空
	Size() int     // Size 返回链表的长度
}
