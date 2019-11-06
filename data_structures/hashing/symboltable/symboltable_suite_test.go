package st

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "go-algorithms/data_structures/types"
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

		orderedSt OrderedST
		letters   = [...]string{
			"S", "E", "A", "R", "C", "H",
			"E", "X", "A", "M", "P", "L", "E",
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

	AssertOrderedSTBehavior := func() {
		It("should be empty", func() {
			Expect(orderedSt.IsEmpty()).To(BeTrue())
			Expect(orderedSt.Size()).To(BeZero())
		})

		Specify("put letters", func() {
			for i, letter := range letters {
				orderedSt.Put(Key(letter), Value(i))
			}

			Expect(orderedSt.Size()).To(Equal(10))
			Expect(orderedSt.Min()).To(BeEquivalentTo("A"))
			Expect(orderedSt.Max()).To(BeEquivalentTo("X"))
		})

		It("can get value by key", func() {
			for _, key := range orderedSt.Keys() {
				_, err := orderedSt.Get(key)
				Expect(err).ShouldNot(HaveOccurred())
			}
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

	Describe("BinarySearchST", Ordered, func() {
		BeforeAll(func() {
			orderedSt = NewBinarySearchST()
		})

		AssertOrderedSTBehavior()
	})
})

func TestSymbolTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SymbolTable Suite")
}
