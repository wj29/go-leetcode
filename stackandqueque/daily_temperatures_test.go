package stackandqueque

import "testing"

// 739. 每日温度
// https://leetcode-cn.com/problems/daily-temperatures/
func Test_DailyTemperatures(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	t.Log(dailyTemperatures(T))
}

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	ret := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			index := len(stack) - 1
			current := stack[index]
			stack = stack[:index]
			ret[current] = i - current
		}
		stack = append(stack, i)
	}
	return ret
}
