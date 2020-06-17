package leetcode

import (
	"errors"
	"testing"
)

// 在未排序的数组中找到第k大的数字，请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素 medium
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/

// TOP K问题，维护一个最大堆或者最小堆来实现
// 第K大元素，维护一个K最大堆，最大堆需要小顶堆实现，小顶堆表示堆顶元素是堆中最小元素
// 现对K个元素建堆，时间复杂度logK，与剩下的N-K个元素比较，时间负载度(N-K)logK 即NlogK
// 最大堆或者最小堆为完全二叉树，需要小顶堆或者大顶堆实现，父节点小于等于(大于等于)孩子结点，利用数组实现这个数据结构
// 父结点下标i则左右孩子结点下标为2i+1和2i+2

func Test_FindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 6
	//t.Log(findKthLargestByQSort(nums, k))
	t.Log(findKthLargestByHeap(nums, k))
}

// 使数组排序即可得到，快排时间复杂度nlogn
func findKthLargestByQSort(nums []int, k int) int {
	if k > len(nums) || k < 0 {
		panic(errors.New("invalid k"))
	}
	qSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

// 快排
func qSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := nums[start]
	l := start
	r := end
	for {
		for nums[r] >= pivot && r > start {
			r--
		}
		for nums[l] <= pivot && l < end {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		} else {
			break
		}
	}
	nums[start], nums[r] = nums[r], nums[start]
	qSort(nums, start, r)
	qSort(nums, r+1, end)
}

func findKthLargestByHeap(nums []int, k int) int {
	if k > len(nums) || k < 0 {
		panic(errors.New("invalid k"))
	}
	nums = heapSort(nums, k)
	return nums[0]
}

func heapSort(nums []int, k int) []int {
	l := len(nums)
	buildHeap(nums, k) // 构建k个元素的最小堆 堆定元素即为这k个元素中最小的，即第k个
	for i := l - 1; i >= k; i-- {
		// 将剩余的数与堆顶元素比较，小于等于直接继续，大于则交换位置进行排序
		if nums[i] <= nums[0] {
			continue
		}
		swap(nums, 0, i)
		heap(nums, 0, k)
	}
	return nums
}

// 堆排序 最小堆
func buildHeap(nums []int, len int) { // 传入了整个nums，但是排序只排了前K个
	// 完全二叉树最少有一半以上的结点都是叶子结点(等比数列性质)
	// 叶子结点不需要向下置换
	for i := len / 2; i >= 0; i-- { // 从非叶子结点开始排序 从下往上构建
		heap(nums, i, len)
	}
}

// 最小堆，即大顶堆实现，顶点最大
func heap(nums []int, i, len int) {
	l := 2*i + 1
	r := 2*i + 2

	min := i
	if l < len && nums[l] < nums[min] { // l >= len时说明没有左孩子结点
		min = l
	}
	if r < len && nums[r] < nums[min] {
		min = r
	}
	if min != i { // 最小值已经改变
		swap(nums, min, i)
		heap(nums, min, len) // 会影响孩子结点的孩子子树(都需要调整)
	}
}

func swap(nums []int, m, n int) []int {
	nums[m], nums[n] = nums[n], nums[m]
	return nums
}
