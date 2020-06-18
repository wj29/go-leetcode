package sort

// 堆排序是不稳定排序，时间复杂度Nlog2(N)，2为对数底数
func heapSort(arr []int) {
	// 建堆 只能保证堆顶元素是最大(小)
	//  1 3 2 5 4 6 满足建堆且堆顶最大(小)
	for i := len(arr); i >= 0; i-- { // 建堆保证上下的单调性一致
		heap(arr, i, len(arr))
	}
	// 将堆顶排到切片末尾(降序)
	length := len(arr)
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 将堆顶元素防到末尾，剩下的进行堆排序
		length--
		heap(arr, 0, length)
	}
}

func heap(arr []int, index, length int) {
	l := 2*index + 1 // 左孩子结点
	r := 2*index + 2 // 右孩子结点
	min := index

	if l < length && arr[l] < arr[min] { // l >= length表示没有该孩子结点
		min = l
	}
	if r < length && arr[r] < arr[min] {
		min = r
	}
	if min != index {
		arr[min], arr[index] = arr[index], arr[min]
		heap(arr, min, length)
	}
}