package sorts

func MergeSorted(items []int) []int {
	return mergeSorted(items)
}

func mergeSorted(items []int) []int {
	length := len(items)

	mid := length / 2
	if mid == 0 {
		return items
	}

	leftItems := items[:mid]
	rightItems := items[mid:]
	return merged(mergeSorted(leftItems), mergeSorted(rightItems))
}

func merged(leftItems, rightItems []int) []int {
	items := make([]int, len(leftItems)+len(rightItems))

	i, j := 0, 0
	for i < len(leftItems) && j < len(rightItems) {
		if leftItems[i] <= rightItems[j] {
			items[i+j] = leftItems[i]
			i++
		} else {
			items[i+j] = rightItems[j]
			j++
		}
	}

	for ; i < len(leftItems); i++ {
		items[i+j] = leftItems[i]
	}

	for ; j < len(rightItems); j++ {
		items[i+j] = rightItems[j]
	}

	return items
}
