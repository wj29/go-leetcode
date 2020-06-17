package doublepoint

import "testing"

// 给定一个字符串和一个字符串数组，找到数组里面最长的字符串，该字符串可以通过删除给定字符串的某些字符得到，
// 答案不唯一时返回长度最长且顺序最小的字符串，答案不存在返回空字符串 medium
// 输入: s = "abpcplea", d = ["ale","apple","monkey","plea"]
// 输出: "apple"
// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/description/

func Test_FindLongestWord(t *testing.T) {
	s := "bab"
	//d := []string{"ale", "apple", "monkey", "plea"}
	d := []string{"ba", "ab", "a"}
	t.Log(findLongestWord(s, d))
}

// 找到s的所有字串逐一与d比较，寻找所有字串的时间负载的2^n
// 通过比较d的所有值是否为s的字串n*m， n为s的长度，m为d的长度
func findLongestWord(s string, d []string) string {
	r := ""
	if len(s) <= 0 {
		return r
	}
	for i := range d {
		if subString(s, d[i]) {
			if len(d[i]) > len(r) {
				r = d[i]
			}
		}
	}
	return r
}

// 此解法也可以做判断子序列的解法
// https://leetcode-cn.com/problems/is-subsequence/ easy
func subString(s string, sub string) bool {
	if len(sub) <= 0 {
		return true
	}
	p1 := 0
	p2 := 0
	for p1 < len(s) {
		if s[p1] == sub[p2] {
			p1++
			p2++
		} else {
			p1++
		}
		if p2 == len(sub) {
			return true
		}
	}
	return false
}
