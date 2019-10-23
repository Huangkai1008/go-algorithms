package linkedlist

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	dll "go-algorithms/data_structures/linkedlist/doublylinkedlist"
)

var _ = Describe("LinkedList", func() {
	var linkedList LinkedList
	AssertLinkedListBehavior := func() {
		It("should be empty", func() {
			Expect(linkedList.IsEmpty()).To(BeTrue())
			Expect(linkedList.Size()).To(BeZero())
		})

		Specify("append 10 items", func() {
			for i := 0; i < 10; i++ {
				linkedList.Append(i)
			}
			Expect(linkedList.Size()).To(Equal(10))

			By("Print the linkedList")
			GinkgoWriter.Println(linkedList)
		})

		It("knows whether the item contains", func() {
			Expect(linkedList.Contains(0)).To(BeTrue())
			Expect(linkedList.Contains(11)).To(BeFalse())
		})

		Specify("insert item", func() {
			for i := 10; i < 20; i++ {
				linkedList.Append(i)
			}
			linkedList.Insert(2, 99)
		})

		It("can remove element", func() {
			linkedList.Delete(8)
			linkedList.Delete(18)
			linkedList.Delete(3)
			linkedList.Delete(1)
		})

		It("can delete element", func() {
			linkedList.Delete(8)
			linkedList.Delete(18)
			linkedList.Delete(3)
			linkedList.Delete(1)
		})

		Specify("clear all the elements", func() {
			size := linkedList.Size()
			for i := 0; i < size; i++ {
				linkedList.Remove(0)
			}
		})

		It("should not be empty again", func() {
			Expect(linkedList.IsEmpty()).To(BeTrue())

			By("Print the linkedList")
			GinkgoWriter.Println(linkedList.Size())
		})

	}

	Describe("SingleLinkedList", Ordered, func() {
		BeforeAll(func() {
			linkedList = NewSll()
		})

		AssertLinkedListBehavior()
	})

	Describe("CircularLinkedList", Ordered, func() {
		BeforeAll(func() {
			linkedList = NewCll()
		})

		AssertLinkedListBehavior()
	})

	Describe("DoublyLinkedList", Ordered, func() {
		BeforeAll(func() {
			linkedList = dll.NewDll()
		})

		AssertLinkedListBehavior()
	})
})

func TestLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LinkedList Suite")
}
