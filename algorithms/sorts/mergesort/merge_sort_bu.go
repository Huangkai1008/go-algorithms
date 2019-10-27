package sorts

// MergeSortBU 归并排序自底而上实现
func MergeSortBU(items []int) {
	length := len(items)
	aux := make([]int, length)
	copy(aux, items)

	for sz := 1; sz < length; sz = sz + sz {
		for low := 0; low < length-sz; low = low + sz + sz {
			merge(items, low, low+sz-1, minInt(low+sz+sz-1, length-1), aux)
		}
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
