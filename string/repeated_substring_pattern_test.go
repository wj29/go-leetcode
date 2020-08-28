package string

import "testing"

// 459. 重复的子字符串
// https://leetcode-cn.com/problems/repeated-substring-pattern/
func Test_RepeatedSubstringPattern(t *testing.T) {
	t.Log(repeatedSubstringPattern("abaabaa"))
}

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return true
	}
	index := 1
	for index <= len(s)/2 {
		if match(s[:index], s) {
			return true
		}
		index++
	}
	return false
}

func match(nums1, nums2 string) bool {
	for i := 0; i < len(nums2); i += len(nums1) {
		if i+len(nums1) > len(nums2) || nums1 != nums2[i:i+len(nums1)] {
			return false
		}
	}
	return true
}

// https://leetcode-cn.com/problems/repeated-substring-pattern/comments/
// 假设给定字符串s可由一个子串x重复n次构成，即s=nx。 现构造新字符串t=2s，即两个s相加，由于s=nx，则t=2nx。
// 去掉t的开头与结尾两位，则这两处的子串被破坏掉，此时t中包含2n-2个子串。
// 由于t中包含2n-2个子串，s中包含n个子串，若t中包含s，则有2n-2>=n，可得n>=2，由此我们可知字符串s可由一个子串x重复至少2次构成，判定为true；
// 反之，若t中不包含s，则有2n-2<n，可得n<2，n只能为1，由此我们可知字符串s=x，假定的子串就为s本身，判定为false。