package backtracking

import (
	"strconv"
	"strings"
	"testing"
)

// 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
// 有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。
// https://leetcode-cn.com/problems/restore-ip-addresses
func Test_RestoreIpAddresses(t *testing.T) {
	s := "25525511135"
	t.Log(restoreIpAddresses(s))
}

// 回溯法尝试所有可能
func restoreIpAddresses(s string) []string {
	ret := make([]string, 0)
	backTraceForIp(s, "", &ret, 0)
	return ret
}

func backTraceForIp(s string, prev string, ret *[]string, start int) {
	if len(prev)-4 == len(s) {
		*ret = append(*ret, prev[:len(prev)-1])
		return
	}
	for i := start; i < len(s); i++ {
		if i-start > 2 { // 每一部分不能超过3位数
			return
		}
		if len(s[start:i+1]) > 1 && s[start] == '0' { // 不合法的0开头的正数
			continue
		}
		if v, _ := strconv.Atoi(s[start : i+1]); v > 255 { // 超过255的整数
			continue
		}
		cnt := len(strings.Split(prev, ".")) - 1
		if len(s)-i > 3*(4-cnt) { // 剩余的长度超过需要的剩余ip部分的最大长度
			continue
		}
		backTraceForIp(s, prev+s[start:i+1]+".", ret, i+1)
	}
}
