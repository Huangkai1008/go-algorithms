package array

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array", func() {
	var array Array
	AssertArrayBehavior := func() {
		It("should be empty", func() {
			Expect(array.IsEmpty()).To(BeTrue())
			Expect(array.Size()).To(BeZero())
		})

		Specify("append 19 items", func() {
			for i := 0; i < 19; i++ {
				array.Append(i)
			}

			Expect(array.Capacity()).To(Equal(20))
			Expect(array.Size()).To(Equal(19))
		})

		Specify("set item", func() {
			array.Set(0, 1)
			Expect(array.Get(0)).To(Equal(1))
		})

		Specify("insert item", func() {
			array.Insert(18, 99)
		})

		It("knows whether the item contains", func() {
			Expect(array.Search(99)).To(Equal(18))
			Expect(array.Search(100)).To(Equal(-1))
			Expect(array.Contains(99)).To(BeTrue())
			Expect(array.Contains(100)).To(BeFalse())
		})

		It("can remove element", func() {
			array.Remove(0)
			array.Remove(1)
		})

		It("should not be empty", func() {
			Expect(array.IsEmpty()).To(BeFalse())
			Expect(array.Size() > 0).To(BeTrue())

			By("Print the array")
			GinkgoWriter.Println(array)
		})

	}

	Describe("DynamicArray", Ordered, func() {
		Context("when creates with default options", func() {
			BeforeAll(func() {
				array = New()
			})

			AssertArrayBehavior()
		})

		Context("when creates with capacity", func() {
			BeforeAll(func() {
				array = New(WithCapacity(20))
			})

			AssertArrayBehavior()
		})
	})
})

func TestArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Array Suite")
}
