package dynamicprogramming

import (
	"math"
	"sort"
	"testing"
)

// 股票系列
// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格
// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票)

func Test_maxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	t.Log(maxProfit1_1(prices))
	t.Log(maxProfit1_2(prices))
	t.Log(maxProfit2(prices))
	t.Log(maxProfit3(prices))
}

// 买卖股票1:
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/ easy

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

// 买卖股票2:
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii easy

// 买过股票在下一个跌之前将股票卖掉，下一个上涨前将股票买入
// 转换上面的思路，每个上升的差值等于每一小段的差值
func maxProfit2(prices []int) int {
	out := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			out += prices[i+1] - prices[i]
		}
	}
	return out
}

// 买卖股票3:
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii hard

// 遍历数组，分别计算左边和右边数组的最大利润，拆分成两个股票I的问题
// 时间复杂度N^2
// 解法不通用，最大买卖次数变成三次即失效
func maxProfit3(prices []int) int {
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

// 买卖股票4:
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/ hard

// 买卖股票含冷冻期
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown

// 买卖股票含手续费
// 给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
// 返回获得利润的最大值。
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
