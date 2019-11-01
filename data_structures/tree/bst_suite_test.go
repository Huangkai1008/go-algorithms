package tree

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BST", func() {
	var bst *BST
	AssertBSTBehavior := func() {
		It("should be empty", func() {
			Expect(bst.IsEmpty()).To(BeTrue())
			Expect(bst.Size()).To(BeZero())
		})

		It("can Add nodes", func() {
			bst.Add(10)
			bst.Add(19)
			bst.Add(5)
			bst.Add(18)
			bst.Add(27)
			bst.Add(1)
			bst.Add(10)
			bst.Add(18)
			bst.Add(15)
			bst.Add(22)

			Expect(bst.IsEmpty()).To(BeFalse())
			Expect(bst.Size()).To(Equal(8))
		})

		It("can know whether node contains", func() {
			Expect(bst.Contains(10)).To(BeTrue())
			Expect(bst.Contains(19)).To(BeTrue())
			Expect(bst.Contains(5)).To(BeTrue())
			Expect(bst.Contains(11)).To(BeFalse())
		})

		Specify("pre order traverse", func() {
			bst.PreOrder()
			GinkgoWriter.Println("\n")
		})

		Specify("in order traverse", func() {
			bst.InOrder()
			GinkgoWriter.Println("\n")
		})

		Specify("post order traverse", func() {
			bst.PostOrder()
			GinkgoWriter.Println("\n")
		})

		Specify("pre order nr traverse", func() {
			bst.PreOrderNR()
			GinkgoWriter.Println("\n")
		})

		Specify("level order", func() {
			bst.LevelOrder()
			GinkgoWriter.Println("\n")
		})
	}

	Describe("BST", Ordered, func() {
		BeforeAll(func() {
			bst = NewBST()
		})

		AssertBSTBehavior()
	})
})

func TestBST(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BST Suite")
}
