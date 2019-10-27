package sorts

// MergeSort 归并排序自顶而下实现
//
// References: https://en.wikipedia.org/wiki/Merge_sort
func MergeSort(items []int) {
	aux := make([]int, len(items))
	mergeSort(items, 0, len(items)-1, aux)
}

func mergeSort(items []int, low, high int, aux []int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2
	mergeSort(items, low, mid, aux)
	mergeSort(items, mid+1, high, aux)
	merge(items, low, mid, high, aux)
}

func merge(items []int, low, mid, high int, aux []int) {
	copy(aux[low:high+1], items[low:high+1])

	for i, j, k := low, mid+1, low; k <= high; k++ {
		if i > mid {
			items[k] = aux[j]
			j++
		} else if j > high {
			items[k] = aux[i]
			i++
		} else if aux[i] <= aux[j] {
			items[k] = aux[i]
			i++
		} else {
			items[k] = aux[j]
			j++
		}
	}
}
