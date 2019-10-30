package sorts

// BubbleSort 冒泡排序
func BubbleSort(items []int) {
	length := len(items)

	for i := 0; i < length-1; i++ {
		swapped := false
		for j := 0; j < length-1-i; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
