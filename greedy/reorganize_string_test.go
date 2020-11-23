package greedy

import (
    "sort"
    "testing"
)

// 767. 重构字符串
// 给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
// 若可行，输出任意可行的结果。若不可行，返回空字符串
// https://leetcode-cn.com/problems/reorganize-string/ medium
func Test_ReorganizeString(t *testing.T) {
    S := "abbabbaaabc"
    t.Log(reorganizeString(S))
    S = "aaab"
    S = "aab"
    t.Log(reorganizeString2(S))
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

// 只需要找到一种解法将两个不相同的字符隔开即可
func reorganizeString2(S string) string {
    m := make(map[int]int) // 存储每个字符的个数
    for _, v := range S {
        m[int(v)-'a'] += 100
    }
    arr := make([]int, 26)
    for i := range arr {
        arr[i] = m[i] + i // 字符信息加载到数组中
    }
    // 比如arr[1] = 201，即代表了b = 'b' - 'a' + 200
    // 通过arr的值可以直接得到该字符arr[i]%100+'a'和出现的次数arr[i]/100
    sort.Ints(arr)
    if arr[25]/100 > (len(S)+1)/2 { // 最大的字符次数超过了半数以上
        return ""
    }
    ret := make([]byte, len(S))
    pos := 0
    for i := 25; i >= 0; i-- { // 优先安排次数多的字符
        tmp := arr[i]%100 + 'a' // 取出字符
        count := arr[i] / 100
        for count > 0 {
            if pos >= len(S) {
                pos = 1
            }
            ret[pos] = byte(tmp)
            pos += 2
            count--
        }
    }
    return string(ret)
}
