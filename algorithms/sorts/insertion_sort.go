package sorts

// InsertionSort 插入排序
//
// 对于随机排序的长度为N且主键不重复的数组：
//     - 平均情况下插入排序需要~N^2/4次比较和~N^2/4次交换
//     - 最坏情况下插入排序需要~N^2/2次比较和~N^2/2次交换
//     - 最好情况下插入排序需要N-1次比较和0次交换
//
// Ref: https://en.wikipedia.org/wiki/Insertion_sort
//
// 假设列表 items`的长度为`N`
// 每次外层循环开始时:
// 		items[0...i)为已排序好的区间，items[i...N)为需要排序的区间
// 每次内层循环结束时(即内层循环全部完成后):
//      items[i]放到items[0...i]中合理的位置
//
func InsertionSort(items []int) {
	for i := 1; i < len(items); i++ {
		insertItem, j := items[i], i-1
		for ; j >= 0 && insertItem < items[j]; j-- {
			items[j+1] = items[j]
		}
		items[j+1] = insertItem
	}
}
