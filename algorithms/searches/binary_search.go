package searches

// BinarySearchByRecur 二分查找法递归实现
//
// 二分查找算法(binary search)又称折半搜索(half-interval search)
// 用于在有序数组中查找特定元素的搜索算法
// 注意给定数组必须是升序排列的
//
// References: https://en.wikipedia.org/wiki/Binary_search_algorithm
func BinarySearchByRecur(items []int, target int) int {
	return binarySearchByRecur(items, target, 0, len(items)-1)
}

func binarySearchByRecur(items []int, target, low, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2
	if items[mid] == target {
		return mid
	} else if items[mid] > target {
		return binarySearchByRecur(items, target, low, mid-1)
	} else {
		return binarySearchByRecur(items, target, mid+1, high)
	}
}
