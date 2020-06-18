package sort

// 插入排序是稳定排序，时间复杂度N^2，空间复杂度1
// 由于插入排序是将后序的数据插入前面已经有序的数据中，从后往前遍历，未排序的相同关键字会在后面，保证了相对顺序
func insertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		backup := arr[i]
		j := i - 1
		//将选出的被排数比较后插入左边有序区
		for j >= 0 && arr[j] > backup { //注意j >= 0必须在前边，否则会数组越界
			arr[j+1] = arr[j] //移动有序数组
			j--               //反向移动下标
		}
		arr[j+1] = backup //插入移动后的空位
	}
}