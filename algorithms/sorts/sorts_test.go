package sorts

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"

	sorts "go-algorithms/algorithms/sorts/mergesort"
)

var _ = Describe("Sorts", func() {

	var (
		items      []int
		experiment *gmeasure.Experiment
	)

	createUnsortedArray := func() {
		By("Create a unsorted array")
		items = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

		By("Create experiment")
		experiment = gmeasure.NewExperiment("Sort Measurement")
		AddReportEntry(experiment.Name, experiment)
	}

	AssertSortBehavior := func(sort Sort) {
		BeforeEach(createUnsortedArray)

		It("should be sorted", func() {
			expected := []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

			By("Sort array one times")
			sort(items)
			Expect(items).To(Equal(expected))
		})

		It("sort array 1k times", Serial, Label("measurement"), func() {
			experiment.SampleDuration("Benchmark1K", func(idx int) {
				sort(items)
			}, gmeasure.SamplingConfig{N: 1 << 10}, gmeasure.Precision(time.Nanosecond))
		})

		It("sort array 64k times", Serial, Label("measurement"), func() {
			experiment.SampleDuration("Benchmark64K", func(idx int) {
				sort(items)
			}, gmeasure.SamplingConfig{N: 1 << 16}, gmeasure.Precision(time.Nanosecond))
		})
	}

	AssertSortedBehavior := func(sort Sorted) {
		BeforeEach(createUnsortedArray)

		It("should be sorted", func() {
			expected := []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

			By("Sort array one times")
			Expect(sort(items)).To(Equal(expected))
		})

		It("sort array 1k times", Serial, Label("measurement"), func() {
			experiment.SampleDuration("Benchmark1K", func(idx int) {
				sort(items)
			}, gmeasure.SamplingConfig{N: 1 << 10}, gmeasure.Precision(time.Nanosecond))
		})

		It("sort array 64k times", Serial, Label("measurement"), func() {
			experiment.SampleDuration("Benchmark64K", func(idx int) {
				sort(items)
			}, gmeasure.SamplingConfig{N: 1 << 16}, gmeasure.Precision(time.Nanosecond))
		})
	}

	Describe("Selection Sort", func() {
		AssertSortBehavior(SelectionSort)
	})

	Describe("Insertion Sort", func() {
		AssertSortBehavior(InsertionSort)
	})

	Describe("Merge Sort", func() {
		AssertSortBehavior(sorts.MergeSort)
	})

	Describe("Merge Sorted", func() {
		AssertSortedBehavior(sorts.MergeSorted)
	})

	Describe("Merge Sort BU", func() {
		AssertSortBehavior(sorts.MergeSortBU)
	})

})

func TestSorts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sorts Suite")
}
