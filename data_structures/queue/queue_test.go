package queue

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", Ordered, func() {
	var queue Queue
	AssertQueueBehavior := func() {
		It("should be empty", func() {
			Expect(queue.IsEmpty()).To(BeTrue())
			Expect(queue.Size()).To(BeZero())
		})

		Specify("Empty queue can't get front or dequeue", func() {
			defer func() {
				if r := recover(); r != nil {
				}
			}()
			Expect(queue.Dequeue()).Should(PanicWith(Equal("Can't dequeue from empty queue.")))
			Expect(queue.GetFront()).Should(PanicWith(Equal("Can't get front from empty queue.")))
		})

		Specify("Enqueue two elements", func() {
			queue.Enqueue(1)
			queue.Enqueue(2)
		})

		It("should not be empty", func() {
			Expect(queue.IsEmpty()).To(BeFalse())
			Expect(queue.Size()).To(Equal(2))
		})

		Specify("The dequeue element should be 1", func() {
			Expect(queue.Dequeue()).To(Equal(1))
		})
		Specify("The front element should be 2 after the 1 dequeued", func() {
			Expect(queue.GetFront()).To(Equal(2))
		})
	}

	Describe("ArrayQueue", func() {
		BeforeAll(func() {
			queue = NewAQ()
		})

		AssertQueueBehavior()
	})

	Describe("LinkedQueue", func() {
		BeforeAll(func() {
			queue = NewLQ()
		})

		AssertQueueBehavior()
	})
})

func TestQueues(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queues Suite")
}
