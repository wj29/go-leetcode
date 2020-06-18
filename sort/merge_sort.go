package sort

// 归并排序是稳定排序,时间复杂度Nlog2(N)，2代表底数，空间复杂度N
// 稳定性的保证在归并时按照原有顺序进行归并即可
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var leftIndex, rightIndex int
	var sortedPile []int
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			sortedPile = append(sortedPile, left[leftIndex])
			leftIndex++
		} else if left[leftIndex] > right[rightIndex] {
			sortedPile = append(sortedPile, right[rightIndex])
			rightIndex++
		} else {
			sortedPile = append(sortedPile, left[leftIndex]) // 此处按原有顺序归并，保证稳定性
			leftIndex++
			sortedPile = append(sortedPile, right[rightIndex])
			rightIndex++
		}
	}

	for leftIndex < len(left) {
		sortedPile = append(sortedPile, left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right) {
		sortedPile = append(sortedPile, right[rightIndex])
		rightIndex++
	}
	return sortedPile
}