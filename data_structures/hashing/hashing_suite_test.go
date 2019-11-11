package hashing

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "go-algorithms/data_structures/types"
)

var _ = Describe("HashTable", func() {
	var (
		hashTable      HashTable
		asciiUppercase = [...]string{
			"A", "B", "C", "D", "E", "F", "G",
			"H", "I", "J", "K", "L", "M", "N",
			"O", "P", "Q", "R", "S", "T", "U",
			"V", "W", "X", "Y", "Z",
		}
	)

	AssertHashTableBehavior := func() {
		It("should be empty", func() {
			Expect(hashTable.IsEmpty()).To(BeTrue())
			Expect(hashTable.Size()).To(BeZero())
		})

		Specify("put letters", func() {
			for i := 0; i < 1000; i++ {
				for _, letter := range asciiUppercase {
					key := fmt.Sprintf("%s%s%d", letter, letter, i)
					hashTable.Put(Key(key), Value(i))
				}
			}
		})

		It("can get value by key", func() {
			value, err := hashTable.Get("AA1")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(value).To(BeEquivalentTo(1))

			value, err = hashTable.Get("ZZ26")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(value).To(BeEquivalentTo(26))

			value, err = hashTable.Get("ZZ")
			Expect(value).To(BeZero())
			Expect(err).Should(MatchError("keyError"))
		})

		It("can delete key and value pair", func() {
			hashTable.Delete("AA1")

			value, err := hashTable.Get("A")
			Expect(value).To(BeZero())
			Expect(err).Should(MatchError("keyError"))
		})

		It("can know whether key contains", func() {
			Expect(hashTable.Contains("A")).To(BeFalse())
			Expect(hashTable.Contains("ZZ1")).To(BeTrue())
		})
	}

	Describe("SeqSearchST", Ordered, func() {
		BeforeAll(func() {
			hashTable = NewSCHashTable(
				WithCapacity(DefaultCapacity),
				WithUpperLoadFactor(DefaultUpperLoadFactor),
				WithLowerLoadFactor(DefaultLowerLoadFactor),
			)
		})

		AssertHashTableBehavior()
	})

	Describe("LinearProbingST", Ordered, func() {
		BeforeAll(func() {
			hashTable = NewLinearProbingHashTable(DefaultLPInitCapacity)
		})

		AssertHashTableBehavior()
	})
})

func TestHashTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HashTable Suite")
}
