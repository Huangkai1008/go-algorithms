package set

import (
	"go-algorithms/data_structures/tree"
	. "go-algorithms/data_structures/types"
)

// BSTSet 通过二分搜索树实现集合
type BSTSet struct {
	bst *tree.BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{
		bst: tree.NewBST(),
	}
}

func (s *BSTSet) Add(data T) {
	s.bst.Add(data)
}

func (s *BSTSet) Remove(data T) {
	s.bst.Remove(data)
}

func (s *BSTSet) Contains(data T) bool {
	return s.bst.Contains(data)
}

func (s *BSTSet) IsEmpty() bool {
	return s.bst.IsEmpty()
}

func (s *BSTSet) Size() int {
	return s.bst.Size()
}
