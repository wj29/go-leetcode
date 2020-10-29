package search

import "testing"

// 131. 分割回文串
// https://leetcode-cn.com/problems/palindrome-partitioning/
// 132. 分割回文串 II
// https://leetcode-cn.com/problems/palindrome-partitioning-ii/
func Test_Partition(t *testing.T) {
    s := "aab"
    t.Log(partition(s))
}

// 深度优先搜索dfs
func partition(s string) [][]string {
    m := make(map[string]bool)
    ret := make([][]string, 0)
    dfsForPartition(s, 0, []string{}, &ret, m)
    return ret
}

func dfsForPartition(s string, start int, prev []string, ret *[][]string, m map[string]bool) {
    if start == len(s) {
        tmp := make([]string, len(prev))
        copy(tmp, prev)
        *ret = append(*ret, tmp)
        return
    }
    for i := start; i < len(s); i++ {
        if !isPalindrome(s, start, i, m) {
            continue
        }
        m[s[start:i+1]] = true
        prev = append(prev, s[start:i+1])
        dfsForPartition(s, i+1, prev, ret, m)
        prev = prev[:len(prev)-1]
    }
}

func isPalindrome(s string, start, end int, m map[string]bool) bool {
    if m[s[start:end+1]] {
        return true
    }
    for start < end {
        if s[start] != s[end] {
            return false
        }
        start++
        end--
    }
    return true
}

// 超时
func minCut(s string) int {
    m := make(map[string]bool)
    ret := len(s)
    dfsForMinCut(s, 0, []string{}, &ret, m)
    return ret
}

func dfsForMinCut(s string, start int, prev []string, ret *int, m map[string]bool) {
    if len(prev) > *ret {
        return
    }
    if start == len(s) {
        *ret = len(prev) - 1
        return
    }
    for i := len(s) - 1; i >= start; i-- {
        if !isPalindrome(s, start, i, m) {
            continue
        }
        m[s[start:i+1]] = true
        prev = append(prev, s[start:i+1])
        dfsForMinCut(s, i+1, prev, ret, m)
        prev = prev[:len(prev)-1]
    }
}

// dp求解
// dp[i]表示s的0-i每个部分为回文可切割的最小次数
// 假设s[:i]是回文，那么dp[i]=0，即不需要切割
// 假设s[:i]不是回文，那么dp[i]与dp[j](j<i)
func minCut2(s string) int {
    return 0
}