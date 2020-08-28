package string

import "testing"

// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列
// https://leetcode-cn.com/problems/zigzag-conversion/
func Test_Convert(t *testing.T) {
	t.Log(convert("ABC", 2))
}

func convert(s string, numRows int) string {
	if numRows < 2 || len(s) <= numRows {
		return s
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		index := i
		for index < len(s) {
			ret += string(s[index])
			if i != 0 && i != numRows-1 {
				if index+(numRows-1-i)*2 < len(s) {
					ret += string(s[index+(numRows-1-i)*2])
				}
			}
			index += 2*numRows - 2
		}
	}
	return ret
}
