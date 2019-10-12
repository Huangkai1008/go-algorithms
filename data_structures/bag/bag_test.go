package bag

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bag", func() {
	var bag *Bag

	BeforeEach(func() {
		By("Create a new bag")
		bag = New()
	})

	Describe("This is a new bag", func() {
		It("should be empty", func() {
			Expect(bag.IsEmpty()).To(BeTrue())
		})

		Specify("The empty bag's size should be zero", func() {
			Expect(bag.Size()).To(BeZero())
		})
	})

	Describe("The bag can add items", func() {
		JustBeforeEach(func() {
			By("Add three items to the bag")
			bag.Add(1)
			bag.Add(2)
			bag.Add(3)
		})

		It("should not be empty", func() {
			Expect(bag.IsEmpty()).To(BeFalse())
		})
		Specify("The bag's size should be three now", func() {
			Expect(bag.Size()).To(Equal(3))
		})
	})
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bag Suite")
}
