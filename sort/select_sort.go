package sort

// 选择排序是一种不稳定排序，时间复杂度N^2，空间复杂度1
// 例如 5 8 5 1 9 在第一次排序后成为 1 8 5 5 9，两个5的顺序交换了
// 类似冒泡排序，通过两层嵌套的循环来遍历、比较数组中元素
// 不同的是选择排序在内层循环中不交换，而是选出最小值的位置，然后在外层循环中交换最小值到头部
// 最终实现数组依次按照最小值排序
// 最好和最坏的时间复杂度都是N^2
func selectSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
	return
}