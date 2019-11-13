package set

import (
	"go-algorithms/algorithms/hashes"
	. "go-algorithms/data_structures/types"
)

const (
	DefaultHashSetCapacity        = 20
	DefaultHashSetUpperLoadFactor = 0.75
	DefaultHashSetLowerLoadFactor = 0.10
)

type Node struct {
	data T
	next *Node
}

type Bucket struct {
	dummyHead *Node
}

// Add 添加元素到bucket首部
//
// 当bucket中已经存在此元素时就不添加
// 返回是否新插入元素
func (b *Bucket) Add(data T) bool {
	if !b.Contains(data) {
		node := &Node{
			data: data,
			next: b.dummyHead.next,
		}
		b.dummyHead.next = node
		return true
	}
	return false
}

// Delete 删除指定元素
//
// 当bucket中不存在此元素时就不删除
// 返回是否删除此元素
func (b *Bucket) Delete(data T) bool {
	for pre, cur := b.dummyHead, b.dummyHead.next; cur != nil; pre, cur = cur, cur.next {
		if cur.data == data {
			pre.next = cur.next
			cur.next = nil
			return true
		}
	}
	return false
}

// Contains 判断元素是否在bucket中
func (b *Bucket) Contains(data T) bool {
	for cur := b.dummyHead.next; cur != nil; cur = cur.next {
		if cur.data == data {
			return true
		}
	}
	return false
}

// Keys 返回bucket中所有的元素
func (b *Bucket) keys() []T {
	var keys []T
	for cur := b.dummyHead.next; cur != nil; cur = cur.next {
		keys = append(keys, cur.data)
	}
	return keys
}

type Option func(*HashSet)

// HashSet 哈希集合表
type HashSet struct {
	buckets         []*Bucket // 集合的存储链表列表
	initCapacity    int64     // 集合的初始容量
	capacity        int64     // 集合的容量
	size            int64     // 集合中的元素个数
	upperLoadFactor float64   // 集合最大负载因子
	lowerLoadFactor float64   // 集合最小负载因子
}

// NewHashSet 返回一个新的哈希集合表
func NewHashSet(opts ...Option) *HashSet {
	set := HashSet{
		capacity:        DefaultHashSetCapacity,
		initCapacity:    DefaultHashSetCapacity,
		upperLoadFactor: DefaultHashSetUpperLoadFactor,
		lowerLoadFactor: DefaultHashSetLowerLoadFactor,
	}
	for _, o := range opts {
		o(&set)
	}

	set.buckets = make([]*Bucket, set.capacity)
	for idx := range set.buckets {
		set.buckets[idx] = &Bucket{
			&Node{},
		}
	}

	return &set
}

func NewHashSetWithCapacity(capacity int64) Option {
	return func(s *HashSet) {
		s.capacity = capacity
		s.initCapacity = capacity
	}
}

func NewHashSetWithUpperLoadFactor(upperLoadFactor float64) Option {
	return func(s *HashSet) {
		s.upperLoadFactor = upperLoadFactor
	}
}

func NewHashSetWithLowerLoadFactor(lowerLoadFactor float64) Option {
	return func(s *HashSet) {
		s.lowerLoadFactor = lowerLoadFactor
	}
}

func (s *HashSet) Add(data T) {
	if s.loadFactor() >= s.upperLoadFactor {
		s.resize(s.capacity * 2)
	}

	hashcode := s.hash(data)
	added := s.buckets[hashcode].Add(data)
	if added {
		s.size++
	}
}

func (s *HashSet) Remove(data T) {
	hashcode := s.hash(data)
	deleted := s.buckets[hashcode].Delete(data)
	if deleted {
		s.size--

		if s.loadFactor() <= s.lowerLoadFactor && s.capacity>>2 >= s.initCapacity {
			s.resize(s.capacity >> 2)
		}
	}
}

func (s *HashSet) Contains(data T) bool {
	hashcode := s.hash(data)
	return s.buckets[hashcode].Contains(data)
}

func (s *HashSet) IsEmpty() bool {
	return s.size == 0
}

func (s *HashSet) Size() int {
	return int(s.size)
}

// hash 集合的哈希函数
func (s *HashSet) hash(data T) uint64 {
	return hashes.SimpleHash(string(rune(data)), uint64(s.capacity))
}

// loadFactor 集合当前的负载因子
func (s *HashSet) loadFactor() float64 {
	return float64(s.size) / float64(s.capacity)
}

// resize 集合扩缩容
func (s *HashSet) resize(capacity int64) {
	hashset := NewHashSet(NewHashSetWithCapacity(capacity))
	for idx := range s.buckets {
		for _, key := range s.buckets[idx].keys() {
			hashset.Add(key)
		}
	}
	s.buckets = hashset.buckets
	s.capacity = capacity
	s.initCapacity = capacity
}
