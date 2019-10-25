package st

// OrderedST 有序符号表的抽象接口
type OrderedST interface {
	SymbolTable

	Min() Key                          // 最小的键
	Max() Key                          // 最大的键
	Floor(key Key) Key                 // 小于等于给定键的最大键
	Ceil(key Key) Key                  // 大于等于给定键的最小键
	Rank(key Key) int                  // 小于给定键的数量
	Select(k int) Key                  // 排名为K的键
	DeleteMin()                        // 删除最小的键
	DeleteMax()                        // 删除最大的键
	RangeSize(low Key, high Key)       // [low...high]之间键的数量
	RangeKeys(low Key, high Key) []Key // [low...high]之间的所有排序后的键
	Keys() []Key                       // 符号表中所有键排序后的列表
}
