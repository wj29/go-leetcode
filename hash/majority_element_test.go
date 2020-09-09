package hash

import "testing"

// 169. 多数元素
// https://leetcode-cn.com/problems/majority-element/
func Test_MajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	t.Log(majorityElement(nums))
}

// 摩尔投票法
// 1.如果候选人不是maj 则 maj,会和其他非候选人一起反对 会反对候选人,所以候选人一定会下台(maj==0时发生换届选举)
// 2.如果候选人是maj, 则maj 会支持自己，其他候选人会反对，同样因为maj 票数超过一半，所以maj 一定会成功当选
// 即从第一个数开始count=1，遇到相同的就加1，遇到不同的就减1，减到0就重新换个数开始计数
func majorityElement(nums []int) int {
	ret := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			ret = nums[i]
			count++
			continue
		}
		if nums[i] == ret {
			count++
		} else {
			count--
		}
	}
	return ret
}
