package deque

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deque", func() {
	var deque Deque
	AssertDequeBehavior := func() {
		It("should be empty", func() {
			Expect(deque.IsEmpty()).To(BeTrue())
			Expect(deque.Size()).To(BeZero())
		})

		Specify("empty deque can't pop or peek", func() {
			defer func() {
				if r := recover(); r != nil {
				}
			}()
			Expect(deque.PopFirst()).Should(PanicWith(Equal("Can't pop from empty deque.")))
			Expect(deque.PopLast()).Should(PanicWith(Equal("Can't pop from empty deque.")))
			Expect(deque.PeekFirst()).Should(PanicWith(Equal("Can't peek from empty deque.")))
			Expect(deque.PeekLast()).Should(PanicWith(Equal("Can't peek from empty deque.")))
		})

		Specify("push 10 elements", func() {
			for i := 1; i <= 5; i++ {
				deque.OfferFirst(i)
			}

			for i := 6; i <= 10; i++ {
				deque.OfferLast(i)
			}

			By("Print the deque:")
			GinkgoWriter.Println(deque)
		})

		It("should not be empty", func() {
			Expect(deque.IsEmpty()).To(BeFalse())
			Expect(deque.Size()).To(Equal(10))
		})

		It("can pop from the front or rear", func() {
			Expect(deque.PopFirst()).To(Equal(5))
			Expect(deque.PopLast()).To(Equal(10))
			Expect(deque.Size()).To(Equal(8))
		})
		It("can peek from the front or rear", func() {
			Expect(deque.PeekFirst()).To(Equal(4))
			Expect(deque.PeekLast()).To(Equal(9))
			By("Print the deque:")
			GinkgoWriter.Println(deque)
		})

	}

	Describe("ArrayDeque", Ordered, func() {
		BeforeAll(func() {
			deque = NewArrayDeque()
		})

		AssertDequeBehavior()
	})

	Describe("LinkedDeque", Ordered, func() {
		BeforeAll(func() {
			deque = NewLinkedDeque()
		})

		AssertDequeBehavior()
	})

	Describe("CircularDeque", Ordered, func() {
		BeforeAll(func() {
			deque = NewCircularDeque(WithCapacity(5))
		})

		AssertDequeBehavior()
	})
})

func TestDeque(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deque Suite")
}
