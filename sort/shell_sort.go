package sort

// shell是不稳定排序，时间复杂度(NlogN)^2，空间复杂度1
// 最佳时间复杂度N，最坏时间复杂度(NlogN)^2
// 插入排序是稳定性排序，但是由于不同的插入排序后最后的稳定性会被打乱
// 小步长增量来插入排序时可以利用大步长增量的有序性，
// 利用了插入排序适合小数据量和基本有序的特点，充分发挥插入排序的优点，
// 使得整体用来处理大数据量的排序非常不错——虽然比快排还差点
func shellSort(arr []int) []int {
	gap := len(arr) / 2

	for gap > 0 {
		for i := 0; i < gap; i++ {
			insertInShell(arr, gap, i)
		}
		gap /= 2
	}
	return arr
}

func insertInShell(arr []int, gap, index int) {
	for i := index + gap; i < len(arr); i += gap {
		backup := arr[i]
		j := i - gap
		for j >= 0 && arr[j] > backup {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
	return
}