package hashing

import (
	"go-algorithms/algorithms/hashes"
	st "go-algorithms/data_structures/hashing/symboltable"
	. "go-algorithms/data_structures/types"
)

type Option func(*SeparateChainingHashTable)

// SeparateChainingHashTable 拉链法实现哈希表
type SeparateChainingHashTable struct {
	buckets         []*st.SequentialSearchST // 存储链表的数组
	initCapacity    int64                    // 哈希表初始化容量
	capacity        int64                    // 哈希表当前容量
	size            int64                    // 哈希表当前键值对个数
	upperLoadFactor float64                  //  哈希表最大负载因子
	lowerLoadFactor float64                  // 哈希表最小负载因子
}

// NewSCHashTable 返回一个新的拉链法实现的哈希表
func NewSCHashTable(opts ...Option) *SeparateChainingHashTable {
	t := SeparateChainingHashTable{
		capacity:        DefaultCapacity,
		upperLoadFactor: DefaultUpperLoadFactor,
		lowerLoadFactor: DefaultLowerLoadFactor,
	}

	for _, o := range opts {
		o(&t)
	}
	t.buckets = make([]*st.SequentialSearchST, t.capacity)
	for idx := range t.buckets {
		t.buckets[idx] = st.NewSequentialSearchST()
	}
	return &t
}

func WithCapacity(capacity int64) Option {
	return func(t *SeparateChainingHashTable) {
		t.capacity = capacity
	}
}

func WithUpperLoadFactor(upperLoadFactor float64) Option {
	return func(t *SeparateChainingHashTable) {
		t.upperLoadFactor = upperLoadFactor
	}
}

func WithLowerLoadFactor(lowerLoadFactor float64) Option {
	return func(t *SeparateChainingHashTable) {
		t.lowerLoadFactor = lowerLoadFactor
	}
}

func (t *SeparateChainingHashTable) Put(key Key, value Value) {
	if t.loadFactor() >= t.upperLoadFactor {
		t.resize(t.capacity * 2)
	}

	hashcode := t.hash(key)
	if !t.buckets[hashcode].Contains(key) {
		t.size++
	}
	t.buckets[hashcode].Put(key, value)
}

func (t *SeparateChainingHashTable) Get(key Key) (Value, error) {
	hashcode := t.hash(key)
	return t.buckets[hashcode].Get(key)
}

func (t *SeparateChainingHashTable) Delete(key Key) {
	hashcode := t.hash(key)
	if t.buckets[hashcode].Contains(key) {
		t.size--
	}
	t.buckets[hashcode].Delete(key)

	if t.capacity > t.initCapacity && t.loadFactor() <= t.lowerLoadFactor {
		t.resize(t.capacity / 2)
	}
}

func (t *SeparateChainingHashTable) Contains(key Key) bool {
	hashcode := t.hash(key)
	return t.buckets[hashcode].Contains(key)
}

func (t *SeparateChainingHashTable) IsEmpty() bool {
	return t.size == 0
}

func (t *SeparateChainingHashTable) Size() int64 {
	return t.size
}

// hash 哈希表的哈希函数
func (t *SeparateChainingHashTable) hash(key Key) uint64 {
	return hashes.SimpleHash(string(key), uint64(t.capacity))
}

// loadFactor 哈希表当前负载因子
// 当前负载因子 = 当前哈希表键值对个数 / 哈希表的当前容量
func (t *SeparateChainingHashTable) loadFactor() float64 {
	return float64(t.size) / float64(t.capacity)
}

// resize 哈希表扩缩容
//
// 创建一个新的扩容容量的哈希表
// 将原有哈希表的所有键值对放入此哈希表中
// 此过程做了重哈希(Rehash)
func (t *SeparateChainingHashTable) resize(chains int64) {
	hashTable := NewSCHashTable(
		WithCapacity(chains),
		WithUpperLoadFactor(t.upperLoadFactor),
		WithLowerLoadFactor(t.lowerLoadFactor),
	)

	for _, bucket := range t.buckets {
		for _, key := range bucket.Keys() {
			value, _ := bucket.Get(key)
			hashTable.Put(key, value)
		}
	}
	t.capacity = chains
	t.size = hashTable.Size()
	t.buckets = hashTable.buckets
}
