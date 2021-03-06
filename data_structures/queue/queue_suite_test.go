package queue

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
	var queue Queue
	AssertQueueBehavior := func() {
		It("should be empty", func() {
			Expect(queue.IsEmpty()).To(BeTrue())
			Expect(queue.Size()).To(BeZero())
		})

		Specify("empty queue can't get front or dequeue", func() {
			defer func() {
				if r := recover(); r != nil {
				}
			}()
			Expect(queue.Dequeue()).Should(PanicWith(Equal("Can't dequeue from empty queue.")))
			Expect(queue.GetFront()).Should(PanicWith(Equal("Can't get front from empty queue.")))
		})

		Specify("enqueue 10 elements", func() {
			for i := 1; i < 11; i++ {
				queue.Enqueue(i)
			}
		})

		It("should not be empty", func() {
			Expect(queue.IsEmpty()).To(BeFalse())
			Expect(queue.Size()).To(Equal(10))
		})

		Specify("the dequeue element should be 1", func() {
			Expect(queue.Dequeue()).To(Equal(1))
		})
		Specify("the front element should be 2 after the 1 dequeued", func() {
			Expect(queue.GetFront()).To(Equal(2))
			GinkgoWriter.Println(queue)
		})

	}

	Describe("ArrayQueue", Ordered, func() {
		BeforeAll(func() {
			queue = NewArrayQueue()
		})

		AssertQueueBehavior()
	})

	Describe("LinkedQueue", Ordered, func() {
		BeforeAll(func() {
			queue = NewLinkedQueue()
		})

		AssertQueueBehavior()
	})

	Describe("CircularQueue", Ordered, func() {
		BeforeAll(func() {
			queue = NewCircularArrayQueue(WithCapacity(5))
		})

		AssertQueueBehavior()
	})
})

func TestQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queue Suite")
}
