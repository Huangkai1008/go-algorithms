package hashes

import "hash/fnv"

// SimpleHash 实现简单的哈希函数
func SimpleHash(key string, capacity uint64) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(key))

	hashValue := h.Sum64()
	return (capacity - 1) & (hashValue ^ (hashValue >> 16))
}
