package doublepoint

import (
    "fmt"
    "testing"
)

// 438. 找到字符串中所有字母异位词
// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
func Test_FindAnagrams(t *testing.T) {
    p := "a"
    s := "a"
    for i := 0; i < 29999; i++ {
        s += "a"
        if i < 10000 {
            p += "a"
        }
    }
    fmt.Print(findAnagrams(s, p))
}

// 不知道为何相同字符串也被认为是字母异位词
func findAnagrams(s string, p string) []int {
    m := make(map[string]bool)
    ret := make([]int, 0)
    for i := 0; i+len(p) <= len(s); i++ {
        if m[s[i:i+len(p)]] {
            ret = append(ret, i)
            continue
        }
        if ectopic(s[i:i+len(p)], p) {
            m[s[i:i+len(p)]] = true
            ret = append(ret, i)
        }
    }
    return ret
}

func ectopic(p1, p2 string) bool {
    m1, m2 := make([]int, 26), make([]int, 26)
    for i := 0; i < len(p1); i++ {
        m1[p1[i]-'a']++
        m2[p2[i]-'a']++
    }
    for i := 0; i < len(m1); i++ {
        if m1[i] != m2[i] {
            return false
        }
    }
    return true
}
