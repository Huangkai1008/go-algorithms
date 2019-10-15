package stack

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	var stack Stack
	AssertQueueBehavior := func() {
		It("should be empty", func() {
			Expect(stack.IsEmpty()).To(BeTrue())
			Expect(stack.Size()).To(BeZero())
		})

		Specify("empty stack can't peek or pop", func() {
			defer func() {
				if r := recover(); r != nil {
				}
			}()
			Expect(stack.Pop()).Should(PanicWith(Equal("Can't pop from empty stack.")))
			Expect(stack.Peek()).Should(PanicWith(Equal("Can't peek from empty stack.")))
		})

		It("push 10 elements", func() {
			for i := 0; i < 10; i++ {
				stack.Push(i)
			}
		})

		It("should not be empty", func() {
			Expect(stack.IsEmpty()).To(BeFalse())
			Expect(stack.Size()).To(Equal(10))
		})

		Specify("the pop element should be 9", func() {
			Expect(stack.Pop()).To(Equal(9))
		})
		Specify("the peek element should be 8 after the 9 popped", func() {
			Expect(stack.Peek()).To(Equal(8))
		})
	}

	Describe("ArrayStack", Ordered, func() {
		BeforeAll(func() {
			stack = NewArrayStack()
		})

		AssertQueueBehavior()
	})

	Describe("LinkedListStack", Ordered, func() {
		BeforeAll(func() {
			stack = NewLLStack()
		})

		AssertQueueBehavior()
	})

	Describe("SliceStack", Ordered, func() {
		BeforeAll(func() {
			stack = NewSliceStack()
		})

		AssertQueueBehavior()
	})
})

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}
