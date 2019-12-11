package heap

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "go-algorithms/data_structures/types"
)

var _ = Describe("Heap", func() {
	var (
		heap Heap
	)

	Describe("Max Heap", Ordered, func() {
		BeforeAll(func() {
			heap = NewMaxHeap()
		})

		It("should be empty", func() {
			Expect(heap.IsEmpty()).To(BeTrue())
			Expect(heap.Size()).To(BeZero())
		})

		Specify("push items to heap", func() {
			for i := 1; i < 11; i++ {
				heap.Push(T(i))
			}
			Expect(heap.IsEmpty()).To(BeFalse())
			Expect(heap.Size()).To(Equal(10))
		})

		Specify("top or pop", func() {
			Expect(heap.Top()).To(BeEquivalentTo(10))
			Expect(heap.Pop()).To(BeEquivalentTo(10))
			Expect(heap.Top()).To(BeEquivalentTo(9))
		})
	})

	Describe("Max Heap With heapify", Ordered, func() {
		BeforeAll(func() {
			heap = NewMaxHeap(
				NewMaxHeapWithItems([]T{3, 8, 1, 2, 9}),
			)
		})

		It("should not be empty", func() {
			Expect(heap.IsEmpty()).To(BeFalse())
			Expect(heap.Size()).To(Equal(5))
		})

		Specify("top or replace", func() {
			Expect(heap.Replace(10)).To(BeEquivalentTo(9))
			Expect(heap.Top()).To(BeEquivalentTo(10))
		})
	})
})

func TestHeap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heap Suite")
}
