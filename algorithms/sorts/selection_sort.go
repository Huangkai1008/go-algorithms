package sorts

// SelectionSort 选择排序
//
// 对于长度为N的数组，选择排序需要大约N^2/2次比较和N次变换
// 选择排序的特点有：
//     1.运行时间和输入无关
//     2.数据移动最少
//
// Ref: https://en.wikipedia.org/wiki/Selection_sort
//
// 假设列表 items`的长度为`N`
// 每次外层循环开始时:
// 		items[0...i)为已排序好的区间，items[i...N)为需要排序的区间
// 每次内层循环结束时(即内层循环全部完成后):
//      items[i...N)中的最小值放到items[i]的位置
//
func SelectionSort(items []int) {
	for i := 0; i < len(items); i++ {
		least := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[least] {
				least = j
			}
		}
		items[i], items[least] = items[least], items[i]
	}
}
