package arrayandmatrix

import (
	"math"
	"testing"
)

// 给定一个二进制数组， 计算其中最大连续1的个数
// https://leetcode-cn.com/problems/max-consecutive-ones/description/  easy
func Test_FindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}
	t.Log(findMaxConsecutiveOnes(nums))
}

// 遇到1则++，通过全局变量取得最大值
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			max = int(math.Max(float64(max), float64(count)))
			count = 0
		}
	}
	// return max // 不能return max, 可能最后一位不是0，最后一个count未比较
	return int(math.Max(float64(max), float64(count)))
}

// 以0分隔字串，最大字串长度
// strings.Split()