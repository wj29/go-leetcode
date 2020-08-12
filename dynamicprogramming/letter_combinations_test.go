package dynamicprogramming

import (
	"testing"
)

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 给出数字到字母的映射与电话按键相同。注意 1 不对应任何字母。
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

func Test_LetterCombinations(t *testing.T) {
	digits := "23"
	t.Log(letterCombinations(digits))
}

// dp[i]表示以第i个数字结尾所能映射的字母种类数组，它是由dp[i-1]加上第i个数字映射的字母得到的
// dp[i] = dp[i-1][j](for j in 0 - len(dp[i-1])) append digits[i]
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[byte][]string{ // create mapping
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	dp := make([][]string, len(digits))
	for i := range dp {
		dp[i] = make([]string, 0)
	}

	dp[0] = append(dp[0], m[digits[0]]...) // base
	for i := 1; i < len(digits); i++ {
		for _, v := range dp[i-1] {
			for _, val := range m[digits[i]] {
				tmp := v + val // can not use v += val, append v, string is a pointer,it will influence the result
				dp[i] = append(dp[i], tmp)
			}
		}
	}
	return dp[len(digits)-1]
}
