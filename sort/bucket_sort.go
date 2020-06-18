package sort

// 桶排序 时间复杂度(N+桶内排序复杂度)，空间复杂度N
// 桶排序的稳定性取决于桶内排序的稳定性
func bucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	num := len(arr)
	max := getMaxInArr(arr)
	bucket := make([][]int, num)
	for i := 0; i < len(arr); i++ {
		index := arr[i] * (num - 1) / max // 所有的arr[i]<=max，最大index为len(num)-1，分块进行排序
		bucket[index] = append(bucket[index], arr[i])
	}
	tmpPos := 0
	for i := 0; i < len(bucket); i++ {
		if len(bucket[i]) > 0 {
			sortImnBucket(bucket[i], 0, len(bucket[i])-1)
			copy(arr[tmpPos:], bucket[i])
			tmpPos += len(bucket[i])
		}
	}
}

func getMaxInArr(arr []int) int {
	if len(arr) <= 0 {
		panic("invalid arr length")
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// 可以选择任意排序算法
func sortImnBucket(arr []int, start, end int) {
	if start >= end {
		return
	}
	left := start
	right := end
	pivot := arr[start]

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
	arr[start], arr[right] = arr[right], arr[start]

	sortImnBucket(arr, start, right)
	sortImnBucket(arr, right+1, end)
}