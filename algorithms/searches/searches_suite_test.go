package searches

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Searches", func() {
	var (
		items      []int
		search     func(items []int, target int) int
		experiment *gmeasure.Experiment
	)

	AssertSortBehavior := func() {
		BeforeEach(func() {
			By("Create a sorted array")
			items = []int{1, 6, 18, 21, 23, 24, 26, 29, 40}

			By("Create experiment")
			experiment = gmeasure.NewExperiment("Sort Measurement")
			AddReportEntry(experiment.Name, experiment)
		})

		It("should can be search", func() {
			By("Search one times")
			Expect(search(items, 26)).To(Equal(6))
			Expect(search(items, 0)).To(Equal(-1))
		})

		It("search 1k times", Serial, Label("measurement"), func() {
			experiment.SampleDuration("Benchmark1K", func(idx int) {
				search(items, 6)
			}, gmeasure.SamplingConfig{N: 1 << 10}, gmeasure.Precision(time.Nanosecond))
		})

		It("search array 64k times", Serial, Label("measurement"), func() {
			experiment.SampleDuration("Benchmark64K", func(idx int) {
				search(items, 6)
			}, gmeasure.SamplingConfig{N: 1 << 16}, gmeasure.Precision(time.Nanosecond))
		})

	}

	Describe("Binary Search With Recursion", func() {
		BeforeEach(func() {
			search = BinarySearchByRecur
		})

		AssertSortBehavior()
	})

	Describe("Binary Search", func() {
		BeforeEach(func() {
			search = BinarySearch
		})

		AssertSortBehavior()
	})
})

func TestSearches(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Searches Suite")
}
