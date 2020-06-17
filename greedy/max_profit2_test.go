package greedy

import (
	"testing"
)

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii easy
func Test_MaxProfit2(t *testing.T) {

}

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
