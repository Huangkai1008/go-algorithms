package array

type Array interface {
	Get(index int) interface{}

	Append(item interface{})

	Set(index int, item interface{})
	Insert(index int, item interface{})

	Remove(index int) interface{}

	Contains(item interface{}) bool
	Search(item interface{}) int

	IsEmpty() bool
	Size() int
	Capacity() int
}

type DynamicArray struct {
	capacity int
	size     int
	items    []interface{}
}

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

func (a *DynamicArray) resize(capacity int) {
	newItems := make([]interface{}, capacity)
	for i := 0; i < a.size; i++ {
		newItems[i] = a.items[i]
	}
	a.capacity = capacity
	a.items = newItems
}
