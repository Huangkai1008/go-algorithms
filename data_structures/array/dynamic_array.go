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

// DynamicArray 实现动态数组
type DynamicArray struct {
	capacity int
	size     int
	items    []interface{}
}

// New 返回一个空动态数组
func New(capacity int) *DynamicArray {
	return &DynamicArray{
		capacity: capacity,
		size:     0,
		items:    make([]interface{}, capacity),
	}
}

func (a *DynamicArray) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get item failed, index out of range.")
	}

	return a.items[index]
}

func (a *DynamicArray) Append(item interface{}) {
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}

	a.items[a.size] = item
	a.size++
}

func (a *DynamicArray) Set(index int, item interface{}) {
	if index < 0 || index >= a.size {
		panic("Set item failed, index out of range.")
	}

	a.items[index] = item
}

func (a *DynamicArray) Insert(index int, item interface{}) {
	if index < 0 || index >= a.size {
		panic("Insert item failed, index out of range.")
	}

	// 当数组的元素数量等于数组容量时需要进行扩容
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}

	for i := a.size - 1; i >= index; i-- {
		a.items[i+1] = a.items[i]
	}
	a.items[index] = item
	a.size++
}

func (a *DynamicArray) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Remove item failed, index out of range.")
	}

	v := a.items[index]
	for i := index; i < a.size-1; i++ {
		a.items[i] = a.items[i+1]
	}
	a.items[a.size-1] = nil
	a.size--

	// 当数组的元素数量等于数组容量的1/4时需要进行缩容
	// 同时需要考虑边界情况，缩容后的容量不能为0
	if a.size == a.capacity/4 && a.capacity/2 != 0 {
		a.resize(a.capacity / 2)
	}

	return v
}

func (a *DynamicArray) Contains(item interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.items[i] == item {
			return true
		}
	}
	return false
}

func (a *DynamicArray) Search(item interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.items[i] == item {
			return i
		}
	}
	return -1
}

func (a *DynamicArray) IsEmpty() bool {
	return a.size == 0
}

func (a *DynamicArray) Size() int {
	return a.size
}

func (a *DynamicArray) Capacity() int {
	return a.capacity
}

// resize 进行数组的扩缩容
func (a *DynamicArray) resize(capacity int) {
	newItems := make([]interface{}, capacity)
	copy(newItems, a.items)
	a.capacity = capacity
	a.items = newItems
}
