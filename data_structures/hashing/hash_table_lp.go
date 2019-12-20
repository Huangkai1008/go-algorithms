package hashing

import (
	"errors"

	"go-algorithms/algorithms/hashes"
	. "go-algorithms/data_structures/types"
)

const (
	DefaultLPInitCapacity    = 4
	DefaultLPUpperLoadFactor = 0.5
	DefaultLPLowerLoadFactor = 0.125
)

// LinearProbingHashTable 线性探测法实现哈希表
type LinearProbingHashTable struct {
	keys            []Key   // 哈希表键的存储数组
	values          []Value // 哈希表值的存储数组
	size            int64   // 哈希表中的键值对数目
	capacity        int64   // 哈希表的容量
	initCapacity    int64   // 哈希表的初始容量
	upperLoadFactor float64 // 哈希表最大负载因子
	lowerLoadFactor float64 // 哈希表最小负载因子
}

// NewLinearProbingHashTable 返回新的哈希表
func NewLinearProbingHashTable(capacity int64) *LinearProbingHashTable {
	return &LinearProbingHashTable{
		keys:            make([]Key, capacity),
		values:          make([]Value, capacity),
		capacity:        capacity,
		initCapacity:    capacity,
		upperLoadFactor: DefaultLPUpperLoadFactor,
		lowerLoadFactor: DefaultLPLowerLoadFactor,
	}
}

func (t *LinearProbingHashTable) Put(key Key, value Value) {
	if t.loadFactor() >= t.upperLoadFactor {
		t.resize(t.capacity * 2)
	}

	hashcode := t.hash(key)
	for ; t.keys[hashcode] != ""; hashcode = (hashcode + 1) % uint64(t.capacity) {
		if t.keys[hashcode] == key {
			t.values[hashcode] = value
		}
	}

	t.keys[hashcode] = key
	t.values[hashcode] = value
	t.size++
}

func (t *LinearProbingHashTable) Get(key Key) (Value, error) {
	for hashcode := t.hash(key); t.keys[hashcode] != ""; hashcode = (hashcode + 1) % uint64(t.capacity) {
		if t.keys[hashcode] == key {
			return t.values[hashcode], nil
		}
	}
	return 0, errors.New("keyError")
}

func (t *LinearProbingHashTable) Delete(key Key) {
	for hashcode := t.hash(key); t.keys[hashcode] != ""; hashcode = (hashcode + 1) % uint64(t.capacity) {
		if t.keys[hashcode] == key {
			t.keys[hashcode] = ""
			t.values[hashcode] = 0

			t.rehash(hashcode)
			t.size--

			// 缩容
			if t.capacity/2 > t.initCapacity && t.loadFactor() <= t.lowerLoadFactor {
				t.resize(t.capacity / 2)
			}
		}
	}
}

func (t *LinearProbingHashTable) Contains(key Key) bool {
	for hashcode := t.hash(key); t.keys[hashcode] != ""; hashcode = (hashcode + 1) % uint64(t.capacity) {
		if t.keys[hashcode] == key {
			return true
		}
	}
	return false
}

func (t *LinearProbingHashTable) IsEmpty() bool {
	return t.size == 0
}

func (t *LinearProbingHashTable) Size() int64 {
	return t.size
}

// hash 哈希表的哈希函数
func (t *LinearProbingHashTable) hash(key Key) uint64 {
	return hashes.SimpleHash(string(key), uint64(t.capacity))
}

// loadFactor 哈希表当前负载因子
// 当前负载因子 = 当前哈希表键值对个数 / 哈希表的当前容量
func (t *LinearProbingHashTable) loadFactor() float64 {
	return float64(t.size) / float64(t.capacity)
}

// resize 哈希表扩缩容
//
// 创建一个新的扩容容量的哈希表
// 将原有哈希表的所有键值对放入此哈希表中
// 此过程做了重哈希(Rehash)
func (t *LinearProbingHashTable) resize(capacity int64) {
	hashTable := NewLinearProbingHashTable(capacity)
	for i := int64(0); i < t.capacity; i++ {
		key := t.keys[i]
		if key != "" {
			hashTable.Put(key, t.values[i])
		}
	}
	t.keys = hashTable.keys
	t.values = hashTable.values
	t.capacity = capacity
	t.initCapacity = capacity
}

// rehash 重哈希
func (t *LinearProbingHashTable) rehash(hashcode uint64) {
	for hashcode = (hashcode + 1) % uint64(t.capacity); t.keys[hashcode] != ""; hashcode = (hashcode + 1) / uint64(t.capacity) {
		key := t.keys[hashcode]
		value := t.values[hashcode]
		t.keys[hashcode] = ""
		t.values[hashcode] = 0
		t.size--

		t.Put(key, value)
	}
}
