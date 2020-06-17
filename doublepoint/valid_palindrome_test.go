package doublepoint

import "testing"

// 给定一个非空字符串，判断是否能通过最多删除一个字符构成回文字符串 easy
// https://leetcode-cn.com/problems/valid-palindrome-ii/description/

func Test_ValidPalindrome(t *testing.T) {
	s := "ebcbbececabbacecbbcbe"
	t.Log(validPalindrome(s))
}

func validPalindrome(s string) bool {
	if len(s) <= 0 {
		return false
	}
	if len(s) <= 2 {
		return true
	}
	b := []byte(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if b[l] != b[r] {
			return f(b, l+1, r) || f(b, l, r-1)
		}
		l++
		r--
	}
	return true
}

func f(b []byte, l, r int) bool {
	for l < r {
		if b[l] != b[r] {
			return false
		}
		l++
		r--
	}
	return true
}
