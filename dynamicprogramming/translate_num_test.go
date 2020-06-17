package dynamicprogramming

import (
	"strconv"
	"testing"
)

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”
// 一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法
// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
func Test_TranslateNum(t *testing.T) {
	num := 12258
	t.Log(translateNum(num))
}

// 对于dp[n]，表示数字n可以翻译的方法数
// 它可以通过翻译最后一位数字或者两位数字得到
// 对于n的最后一位数字如果>5，那么它只可以通过一位数字的方式翻译
// 最后一位数字<=5，可以通过翻译翻译一位数字的方式，也可以通过两位数字的方式翻译(n-1必须是1或2)
// dp[n] = dp[n-1] + dp[n-2](如果n-1和n可以被翻译)
func translateNum(num int) int {
	str := strconv.Itoa(num)
	if len(str) == 1 { // 一位数字只能一种方式
		return 1
	}
	dp := make([]int, len(str)+1)
	dp[0] = 1 // 0或者1不影响，根据dp[2]推算出dp[0]=1
	dp[1] = 1
	prevTwo, _ := strconv.Atoi(str[:2])
	if prevTwo <= 25 {
		dp[2] = 2 // 不存在09这样的数 第一位肯定是1或者2
	} else {
		dp[2] = 1
	}

	two := 0
	for i := 2; i < len(str); i++ {
		two, _ = strconv.Atoi(str[i-1 : i+1])
		dp[i+1] = dp[i]
		if two < 26 && two > 9 { // 在10-25之间时才能被翻译
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(str)]
}
