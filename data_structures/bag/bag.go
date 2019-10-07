package bag

// IBag 是背包的抽象类型
//
// 背包是一种不支持从中删除元素的集合数据结构
type IBag interface {
	Add(item interface{}) // Add 添加元素到背包中
	IsEmpty() bool        // IsEmpty 判断背包是否为空
	Size() int            // Size 返回背包中元素的数量
}

// Bag 是背包的具体实现
type Bag struct {
	items []interface{}
}

// New 返回一个新的背包实例
func New() *Bag {
	return &Bag{}
}

func (b *Bag) Add(item interface{}) {
	b.items = append(b.items, item)
}

func (b *Bag) IsEmpty() bool {
	return len(b.items) == 0
}

func (b *Bag) Size() int {
	return len(b.items)
}
