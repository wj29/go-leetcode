package stackandqueque

import (
    "testing"
)

// 316. 去除重复字母
// https://leetcode-cn.com/problems/remove-duplicate-letters/

// 1081. 不同字符的最小子序列
// https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/
func Test_SmallestSubsequence(t *testing.T) {
    s := "ecbacba" // eacb
    //s = "bcabc"    // abc
    //s = "cbacdcbc" // acdb
    //s = "cdadabcc" // adbc
    //s = "bbcaac"   // bac
    t.Log(smallestSubsequence(s))
}

// 单调栈
// 统计每个字符出现的次数
// 用栈来存储最终返回的字符串，并维持字符串的最小字典序

// 每遇到一个字符，如果这个字符不存在于栈中，就需要将该字符压入栈中
// 若存在栈中不需要考虑(因为当前栈中已经维持了该字符的字典序，即后面的字符会加入而舍弃前面的字符会导致字典序变大)

// 但在压入之前，需要先将之后还会出现(优先选择后面的字符保证较小的字典序)，并且字典序比当前字符小的栈顶字符移除，继续循环，直到栈顶比当前小，然后再将当前字符压入
// 若遇到的的字符之后不会出现(但是他必须出现一次，必须选他)，停止循环，
// 不需要考虑他之前的字符，因为他之前的字符若是比他小且后面有肯定正确，若比他大且后面有，在他入栈时肯定会被弹出，若后面没有了肯定也得选，以此类推前面的字符都是由此得来，故可以停止
func smallestSubsequence(s string) string {
    m := make(map[byte]int)
    for _, v := range []byte(s) {
        m[v]++
    }
    stack := make([]byte, 0)
    inStack := make(map[byte]bool)
    for i := 0; i < len(s); i++ {
        if inStack[s[i]] {
            m[s[i]]--
            continue
        }
        for len(stack) > 0 {
            current := stack[len(stack)-1]
            if m[current] == 0 || current < s[i] {
                break
            }
            stack = stack[:len(stack)-1]
            inStack[current] = false
        }
        stack = append(stack, s[i])
        inStack[s[i]] = true
        m[s[i]]--
    }
    return string(stack)
}
