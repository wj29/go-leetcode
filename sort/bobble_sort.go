package sort

// 冒泡排序是一种稳定排序，平均时间复杂度N^2，空间复杂度1
// 最佳时间复杂度N^2，最坏时间复杂度N^2，无论数据如何排序，复杂度不变
// 对于交换相邻的元素，在满足条件arr[i] > arr[j]或arr[i] < arr[j]时才会发生调换，arr[i] == arr[j]不会调换
func bobbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] { // 选出了最大数
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}