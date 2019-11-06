package st

import (
	"errors"

	. "go-algorithms/data_structures/types"
)

// BinarySearchST 二分查找符号表的实现
//
// 此种实现保证键的存储是有序的
type BinarySearchST struct {
	keys     []Key   // 符号表键的存储数组
	values   []Value // 符号表值的存储数组
	size     int     // 符号表的元素个数
	capacity int     // 符号表的容量
}

// NewBinarySearchST 返回一个新的二分查找符号表
func NewBinarySearchST() *BinarySearchST {
	return &BinarySearchST{
		keys:     make([]Key, DefaultCapacity),
		values:   make([]Value, DefaultCapacity),
		size:     0,
		capacity: DefaultCapacity,
	}
}

func (st *BinarySearchST) Put(key Key, value Value) {
	k := st.Rank(key)

	// 更新键值对
	if k < st.size && st.keys[k] == key {
		st.values[k] = value
		return
	}

	// 插入键值对
	if st.size == st.capacity {
		st.resize(st.capacity * 2)
	}

	for i := st.size - 1; i >= k; i-- {
		st.keys[i+1] = st.keys[i]
		st.values[i+1] = st.values[i]
	}
	st.keys[k] = key
	st.values[k] = value
	st.size++
}

func (st *BinarySearchST) Get(key Key) (Value, error) {
	if st.IsEmpty() {
		return 0, errors.New("keyError")
	}

	k := st.Rank(key)
	if k < st.size && st.keys[k] == key {
		return st.values[k], nil
	} else {
		return 0, errors.New("KeyError")
	}
}

func (st *BinarySearchST) Delete(key Key) {
	if st.IsEmpty() {
		panic("Delete from empty symbol table.")
	}

	k := st.Rank(key)
	if k < st.size && st.keys[k] == key {
		for i := k; i < st.size-1; i++ {
			st.keys[i] = st.keys[i+1]
			st.values[i] = st.values[i+1]
		}

		st.size--

		if st.size == st.capacity/4 && st.capacity/2 != 0 {
			st.resize(st.capacity / 2)
		}
	}

}

func (st *BinarySearchST) Contains(key Key) bool {
	k := st.Rank(key)
	if k < st.size && st.keys[k] == key {
		return true
	}
	return false
}

func (st *BinarySearchST) Min() Key {
	if st.IsEmpty() {
		panic("Get min key from empty symbol table.")
	}

	return st.keys[0]
}

func (st *BinarySearchST) Max() Key {
	if st.IsEmpty() {
		panic("Get max key from empty symbol table.")
	}

	return st.keys[st.size-1]
}

func (st *BinarySearchST) Floor(key Key) Key {
	k := st.Rank(key)

	if k == 0 { // 没有符合条件的键
		return ""
	} else if k < st.size && st.keys[k] == key {
		return st.keys[k]
	} else {
		return st.keys[k-1]
	}

}

func (st *BinarySearchST) Ceil(key Key) Key {
	k := st.Rank(key)

	if k == st.size {
		return ""
	} else {
		return st.keys[k]
	}
}

// Rank 返回小于给定键的数量
//
// 当key小于数组中的所有元素时：返回0(未命中)
// 当key大于数组中的所有元素时：返回N(未命中)
// 当key小于数组中的某个数时：返回这个数的索引(未命中)
// 当命中时：返回key的索引
func (st *BinarySearchST) Rank(key Key) int {
	low, high := 0, st.size-1
	for low <= high {
		mid := low + (high-low)/2
		if st.keys[mid] < key {
			low = mid + 1
		} else if st.keys[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return low
}

func (st *BinarySearchST) Select(k int) Key {
	if k < 0 || k > st.size-1 {
		panic("Select failed, index out of range.")
	}

	return st.keys[k]
}

func (st *BinarySearchST) DeleteMin() {
	st.Delete(st.Min())
}

func (st *BinarySearchST) DeleteMax() {
	st.Delete(st.Max())
}

func (st *BinarySearchST) RangeSize(low Key, high Key) int {
	if low > high {
		return 0
	}

	if st.Contains(high) {
		return st.Rank(high) - st.Rank(low) + 1
	} else {
		return st.Rank(high) - st.Rank(low)
	}
}

func (st *BinarySearchST) RangeKeys(low Key, high Key) []Key {
	if low > high {
		return []Key{}
	}

	startIndex := st.Rank(low)
	endIndex := st.Rank(high)
	if st.Contains(high) {
		endIndex++
	}

	return st.keys[startIndex:endIndex]
}

func (st *BinarySearchST) Keys() []Key {
	return st.keys
}

func (st *BinarySearchST) IsEmpty() bool {
	return st.size == 0
}

func (st *BinarySearchST) Size() int {
	return st.size
}

// resize 二分查找符号表扩缩容
func (st *BinarySearchST) resize(capacity int) {
	newKeys := make([]Key, capacity)
	newValues := make([]Value, capacity)
	copy(newKeys, st.keys)
	copy(newValues, st.values)
	st.keys = newKeys
	st.values = newValues
	st.capacity = capacity
}
