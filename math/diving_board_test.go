package math

import "testing"

// 你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。
// 你必须正好使用k块木板。编写一个方法，生成跳水板所有可能的长度。
// 返回的长度需要从小到大排列。
// https://leetcode-cn.com/problems/diving-board-lcci
func Test_DivingBoard(t *testing.T) {
	shorter, longer, k := 1, 2, 2
	t.Log(divingBoard(shorter, longer, k))
}

// 每新增一块长木板就少一块短木板，增加的长度为longer-shorter
func divingBoard(shorter int, longer int, k int) []int {
	min, max, diff := k*shorter, k*longer, longer-shorter
	ret := make([]int, 0)
	if k == 0 {
		return ret
	}
	if diff == 0 {
		return []int{max}
	}
	for i := min; i <= max; i += diff {
		ret = append(ret, i)
	}
	return ret
}
