package array

// Array 是数组的抽象类型
type Array interface {
	Get(index int) interface{}          // Get 返回数组索引位置元素的值
	Append(item interface{})            // Append 添加元素到数组尾部
	Set(index int, item interface{})    // Set 设置数组索引位置元素的值
	Insert(index int, item interface{}) // Insert 插入元素到数组的索引位置
	Remove(index int) interface{}       // Remove 删除数组索引位置的元素，并返回此元素的值
	Contains(item interface{}) bool     // Contains 判断元素是否在数组中
	Search(item interface{}) int        // Search 返回元素第一次出现的索引位置，如果元素不存在，返回-1

	IsEmpty() bool // IsEmpty 判断数组是否为空
	Size() int     // Size 返回数组中元素的数量
	Capacity() int // Capacity 返回数组的容量
}
