package dynamicprogramming

import (
	"testing"
)

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。
// https://leetcode-cn.com/problems/coin-change medium

// 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个
// https://leetcode-cn.com/problems/coin-change-2/description/  medium
func Test_CoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	t.Log(coinChange(coins, amount))
	//amount = 3
	t.Log(coinChange1(coins, amount))
	amount = 5
	t.Log(change(amount, coins))
}

// 假设dp[amount]为构成amount数额的最少硬币个数
// 零钱的种类有coins种，那么dp[amount]只能由dp[amount-coins(循环)]+1(coins的一种)得到
// 即 dp[amount] = min(dp[amount-coins[i]](for i from 0 - len(coins)) + 1
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = -1
	minCoin := getMin(coins)
	maxCoin := getMax(coins)
	for i := 1; i <= amount; i++ {
		if i < minCoin { // 定义base
			dp[i] = -1
		} else if i > maxCoin { // 先走大于max的可以避免下一个for循环
			tmp := make([]int, 0)
			for _, v := range coins {
				if i-v >= 0 && dp[i-v] > 0 {
					tmp = append(tmp, dp[i-v])
				}
			}
			if len(tmp) > 0 {
				dp[i] = tmp[0] + 1
				for _, p := range tmp {
					if p+1 < dp[i] {
						dp[i] = p + 1
					}
				}
			} else {
				dp[i] = -1
			}
		} else if inCoins(i, coins) {
			dp[i] = 1
		} else {
			tmp := make([]int, 0)
			for _, v := range coins {
				if i-v >= 0 && dp[i-v] > 0 {
					tmp = append(tmp, dp[i-v])
				}
			}
			if len(tmp) > 0 {
				dp[i] = tmp[0] + 1
				for _, p := range tmp {
					if p+1 < dp[i] {
						dp[i] = p + 1
					}
				}
			} else {
				dp[i] = -1
			}
		}
	}
	return dp[amount]
}

func getMin(coins []int) int {
	min := coins[0]
	for _, v := range coins {
		if v < min {
			min = v
		}
	}
	return min
}

func getMax(coins []int) int {
	max := coins[0]
	for _, v := range coins {
		if v > max {
			max = v
		}
	}
	return max
}

func inCoins(val int, coins []int) bool {
	for _, v := range coins {
		if v == val {
			return true
		}
	}
	return false
}

// dp[n]代表凑成金额n所需的最小硬币数
func coinChange1(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	minCoin := getMin(coins)
	for i := 0; i <= amount; i++ {
		// 定义base
		if i < minCoin {
			dp[i] = -1
			continue
		}
		if inCoins(i, coins) {
			dp[i] = 1
			continue
		}
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] > 0 {
				if dp[i] == 0 {
					dp[i] = dp[i-coins[j]] + 1
					continue
				}
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
		if dp[i] == 0 {
			dp[i] = -1
		}
	}
	return dp[amount]
}

// 正确的子问题定义应该是，problem(k,i) = problem(k-1, i) + problem(k, i-k)
// 即前k个硬币凑齐金额i的组合数等于前k-1个硬币凑齐金额i的组合数加上在原来i-k的基础上使用硬币的组合数。
// 说的更加直白一点，那就是用前k的硬币凑齐金额i，要分为两种情况分析，一种是没有用前k-1个硬币就凑齐了，一种是前面已经凑到了i-k， 现在就差第k个硬币了。
// 状态数组就是DP[k][i], 即前k个硬币凑齐金额i的组合数。
// 这里不再是一维数组，而是二维数组。第一个维度用于记录当前组合有没有用到硬币k，第二个维度记录现在凑的金额是多少？如果没有第一个维度信息，当我们凑到金额i的时候，我们不知道之前有没有用到硬币k。
// 因为这是个组合问题，我们不关心硬币使用的顺序，而是硬币有没有被用到。是否使用第k个硬币受到之前情况的影响

// 之前爬楼梯问题中，我们将一维数组降维成点。这里问题能不能也试着降低一个维度，只用一个数组进行表示呢？
// 这个时候，我们就需要重新定义我们的子问题了
// 此时的子问题是，对于硬币从0到k，我们必须使用第k个硬币的时候，当前金额的组合数。
// 因此状态数组DP[i]表示的是对于第k个硬币能凑的组合数

// 题解: https://leetcode-cn.com/problems/coin-change-2/solution/ling-qian-dui-huan-iihe-pa-lou-ti-wen-ti-dao-di-yo/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 什么都不选也是一种方式
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
