package sort

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {
	arr := []int{8, 7, 3, 9, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	//arr = []int{3, 2, 1, 5, 6, 4}
	//qSort(arr, 0, len(arr)-1)
	//insertSort(arr)
	//bucketSort(arr)
	//selectSort(arr)
	//bobbleSort(arr)
	//arr = mergeSort(arr)
	//shellSort(arr)
	heapSort(arr)
	fmt.Println(arr)
}
