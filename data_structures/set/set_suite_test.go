package set

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "go-algorithms/data_structures/types"
)

var _ = Describe("set", func() {
	var set Set
	AssertSetBehavior := func() {
		It("should be empty", func() {
			Expect(set.IsEmpty()).To(BeTrue())
			Expect(set.Size()).To(BeZero())
		})

		Specify("add elements", func() {
			for i := 0; i < 100; i++ {
				set.Add(T(i))
			}

			for i := 0; i < 60; i++ {
				set.Add(T(i))
			}

			Expect(set.Size()).To(Equal(100))
		})

		Specify("remove elements", func() {
			for i := 0; i < 60; i++ {
				set.Remove(T(i))
			}
		})

		It("can know whether key contains", func() {
			Expect(set.Contains(1)).To(Equal(false))
			Expect(set.Contains(100)).To(Equal(false))
			Expect(set.Contains(88)).To(Equal(true))
		})
	}

	Describe("BSTSet", Ordered, func() {
		BeforeAll(func() {
			set = NewBSTSet()
		})

		AssertSetBehavior()
	})

	Describe("HashSet", Ordered, func() {
		BeforeAll(func() {
			set = NewHashSet()
		})

		AssertSetBehavior()
	})

})

func TestSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Set Suite")
}
