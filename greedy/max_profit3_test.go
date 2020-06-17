package greedy

import (
	"fmt"
	"sort"
	"testing"
)

// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii hard
func Test_MaxProfit3(t *testing.T) {
	nums := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	t.Log(maxProfit3_2(nums))
}

// 波峰与波谷差值最大的两个单调递增区间
// 思路错误，只考虑了同一区间的差值，未考虑到不同区间的差值会比同一区间的差值大 []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
func maxProfit3_1(prices []int) int { // wrong
	if len(prices) <= 0 {
		return 0
	}
	f := make([]int, 0)
	f = append(f, prices[0])
	for i := 1; i < len(prices)-1; i++ {
		if (prices[i] <= prices[i+1] && prices[i] <= prices[i-1]) || (prices[i] >= prices[i+1] && prices[i] >= prices[i-1]) {
			f = append(f, prices[i])
		}
	}
	f = append(f, prices[len(prices)-1])
	fmt.Println(f)
	p := make([]int, 0)
	for i := 0; i < len(f)-1; i++ {
		if f[i+1]-f[i] > 0 {
			p = append(p, f[i+1]-f[i])
		}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})
	fmt.Println(p)
	if len(p) <= 0 {
		return 0
	}
	if len(p) <= 1 {
		return p[0]
	}
	return p[0] + p[1]
}

// 遍历数组，分别计算左边和右边数组的最大利润，拆分成两个股票I的问题
// 时间复杂度N^2
// 解法不通用，最大买卖次数变成三次即失效
func maxProfit3_2(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	ret := make([]int, 0)
	for i := 0; i < len(prices); i++ {
		ret = append(ret, maxProfit1_2(prices[:i])+maxProfit1_2(prices[i:]))
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] > ret[j]
	})
	return ret[0]
}
