package string

import "testing"

// 925. 长按键入
// https://leetcode-cn.com/problems/long-pressed-name/
func Test_IsLongPressedName(t *testing.T) {
	name := "alex"
	typed := "aaleex"
	t.Log(isLongPressedName(name, typed))
}

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)

}
