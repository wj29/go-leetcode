package sort

// 快速排序是不稳定排序，时间复杂度Nlog2(N)，2代表底数，空间复杂度1
// 不稳定发生在right和start交换的时候
// 霍尔快排：以最左为基点比较元素,并从两端同时进行比较双向排序, 将双方都需调整位置的两个元素相互交换直到双方相交，
// 然后将最左元素与右端当前锚点交换，最终实现基准点归位。
func qSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := arr[start]
	left := start
	right := end

	for {
		for arr[right] >= pivot && right > start {
			right--
		}
		for arr[left] <= pivot && left < end {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		} else {
			break
		}
	}
	arr[start], arr[right] = arr[right], arr[start] // 会引发不稳定性
	qSort(arr, start, right)
	qSort(arr, right+1, end)
}