package hash

import "testing"

// 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数
// https://leetcode-cn.com/problems/missing-number/ easy
func Test_MissingNumber(t *testing.T) {
	nums := []int{3, 0, 1}
	t.Log(missingNumber(nums))
}

// 排序，找到缺失的那一个 时间复杂度nlogn
// map，将0-n存入map，遍历数组，找到map中没有的那个，时间复杂度n
// 位运算，同一个数字与本身异或得到0，相同为0不同为1，时间复杂度n
// 数学方法0-n的和为n(n+1)/2，减去数组即为缺失的数，时间复杂度n
func missingNumber(nums []int) int {
	missing := len(nums) // 数组缺失一个数，为0到n-1，那么需要加上n
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i] // 利用数组本身下标模拟0-n
	}
	return missing
}
