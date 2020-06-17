package greedy

import (
	"math"
	"testing"
)

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润
// 注意：你不能在买入股票前卖出股票
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/ easy
func Test_MaxProfit1(t *testing.T) {

}

// 最简单方法对数组中每个元素进行计算，得到最小值，时间复杂度1+2+...+n = n(n+1)/2 即N^2
func maxProfit1_1(prices []int) int {
	out := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > out {
				out = prices[j] - prices[i]
			}
		}
	}
	return out
}

// 对于每一天来说，记录它之前的历史最低点，得到当天卖出的利润，遍历整个数组每一天卖出得到的利润最大即为最大利润
// 时间复杂度N
func maxProfit1_2(prices []int) int {
	out := 0
	min := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > out {
			out = prices[i] - min
		}
	}
	return out
}
