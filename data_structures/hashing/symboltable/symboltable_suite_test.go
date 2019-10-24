package st

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SymbolTable", func() {

	var (
		st             SymbolTable
		asciiUppercase = [...]string{
			"A", "B", "C", "D", "E", "F", "G",
			"H", "I", "J", "K", "L", "M", "N",
			"O", "P", "Q", "R", "S", "T", "U",
			"V", "W", "X", "Y", "Z",
		}
	)

	AssertSTBehavior := func() {

		It("should be empty", func() {
			Expect(st.IsEmpty()).To(BeTrue())
			Expect(st.Size()).To(BeZero())
		})

		Specify("put 26 letters", func() {
			for i, letter := range asciiUppercase {
				st.Put(Key(letter), Value(i+1))
			}
		})

		It("can get value by key", func() {
			value, err := st.Get("A")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(value).To(BeEquivalentTo(1))

			value, err = st.Get("Z")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(value).To(BeEquivalentTo(26))

			value, err = st.Get("ZZ")
			Expect(value).To(BeZero())
			Expect(err).Should(MatchError("keyError"))
		})

		It("can delete key and value pair", func() {
			st.Delete("A")
			Expect(st.Size()).To(Equal(len(asciiUppercase) - 1))

			value, err := st.Get("A")
			Expect(value).To(BeZero())
			Expect(err).Should(MatchError("keyError"))
		})

		It("can know whether key contains", func() {
			Expect(st.Contains("A")).To(BeFalse())
			Expect(st.Contains("B")).To(BeTrue())
		})

	}

	Describe("ArrayST", Ordered, func() {
		BeforeAll(func() {
			st = NewArrayST(WithCapacity(5))
		})

		AssertSTBehavior()
	})

	Describe("SeqSearchST", Ordered, func() {
		BeforeAll(func() {
			st = NewSequentialSearchST()
		})

		AssertSTBehavior()
	})
})

func TestSymbolTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SymbolTable Suite")
}
