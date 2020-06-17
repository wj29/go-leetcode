package math

import (
	"testing"
)

// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数
// 进阶: 不将整数转为字符串来解决这个问题
// https://leetcode-cn.com/problems/palindrome-number/  easy
func Test_IsPalindrome(t *testing.T) {
	x := 1000021
	t.Log(isPalindrome(x))
	t.Log(isPalindrome1(x))
}

// 翻转数字判断是否相等
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

// 通过不断除以10得到末尾数字和长度
func isPalindrome1(x int) bool {
	if x >= 0 && x <= 9 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	reverseX, count, tmp := 0, 0, x
	for tmp != 0 { // 获取长度
		tmp /= 10
		count++
	}
	for i := 1; i <= count/2; i++ {
		reverseX = reverseX*10 + (x % 10)
		x /= 10
	}
	if count%2 != 0 {
		x /= 10
	}
	return x == reverseX
}
