package arrayandmatrix

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 5504. 使数组和能被 P 整除
// https://leetcode-cn.com/problems/make-sum-divisible-by-p/
func Test_MinSubarray(t *testing.T) {
	nums := []int{3, 1, 4, 2}
	//nums = []int{6, 3, 5, 2}
	p := 6
	//nums = []int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}
	//p = 148
	t.Log(minSubarray(nums, p))
}

// (a+b)%m = (a%m+b%m)%m
// a*b%m = ((a%m)*(b%m))%m

// 假设nums不能被p整除，记mod=sum(nums)%p
// 即需要找到一个最短的子数组使得nums[i:j]%p=mod，消除多余的p使nums能被p整除
// 记prev[j]为nums[:j]的前缀和，只需要(prev[j]-prev[i])%p=mod
// 根据同余定理，记prev[i]%m = currentMod，需要找到一个目标prev[j]的targetMod使得其前缀和的差为mod
// targetMod = currentMod - mod (+ m)
func minSubarray(nums []int, p int) int {
	ret := len(nums)
	m := make(map[int]int)
	m[0] = -1 // 前缀和-1下标的值
	sum := 0
	mod := (common.Sum(nums)%p + p) % p
	if mod == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		tmpMod := (sum%p + p) % p
		targetMod := tmpMod - mod
		if targetMod < 0 {
			targetMod += p
		}
		if v, ok := m[targetMod]; ok {
			ret = common.Min(i-v, ret)
		}
		m[tmpMod] = i
	}
	if ret == len(nums) {
		return -1
	}
	return ret
}
