package greedy

import (
	"sort"
	"testing"
)

// 给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
// 若可行，输出任意可行的结果。若不可行，返回空字符串
// https://leetcode-cn.com/problems/reorganize-string/ medium
func Test_ReorganizeString(t *testing.T) {
	S := "abbabbaaabc"
	t.Log(reorganizeString(S))
}

// 遍历字符串对字母进行计数操作，用数组进行倒叙排列
// 每次添加个数最多的字母，依次添加后面的字母
// 当做多的字母的个数超过一半以上时则不能成立，返回空
func reorganizeString(S string) string {
	m := make(map[byte]int)
	for i := range S {
		m[S[i]]++
	}
	arr := make([]*cV, len(m))
	for i := range arr {
		arr[i] = new(cV)
	}
	index := 0
	for i, v := range m {
		arr[index].v = i
		arr[index].c = v
		index++
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].c > arr[j].c
	})
	ret := make([]byte, len(S))
	pos := 0
	for _, v := range arr {
		for v.c > 0 {
			v.c--
			if pos >= len(ret) {
				pos = 1
			}
			ret[pos] = v.v
			pos += 2
		}
	}
	for i := 0; i < len(ret)-1; i++ {
		if ret[i] == ret[i+1] {
			return ""
		}
	}
	return string(ret)
}

type cV struct {
	v byte
	c int
}
