package doublepoint

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
// 重复出现的子串要计算它们出现的次数
// https://leetcode-cn.com/problems/count-binary-substrings
func Test_CountBinarySubstrings(t *testing.T) {
	s := "00110011"
	//s = "00110"
	//s = "10101"
	t.Log(countBinarySubstrings(s))
}

// 慢快指针分别指向前一个连续数和下一个连续数的位置
func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	slow, cur, fast, count := 0, 0, 0, 0
	for fast < len(s) {
		for fast < len(s)-1 && s[fast+1] == s[fast] {
			fast++
		}
		fast++
		if cur == fast {
			count++
			break
		}
		count += common.Min(fast-cur, cur-slow)
		slow = cur
		cur = fast
	}
	return count
}
