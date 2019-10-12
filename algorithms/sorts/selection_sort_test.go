package sorts

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Selection sort", func() {
	var items []int

	BeforeEach(func() {
		items = []int{6, 31, 1, 3, 13, 9}
	})

	Describe("Run selection sort", func() {
		It("should be sorted", func() {
			SelectionSort(items)
			Expect(items, []int{1, 3, 6, 9, 13, 31})
		})
	})
})

func TestSelectionSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Selection-Sort Suite")
}
