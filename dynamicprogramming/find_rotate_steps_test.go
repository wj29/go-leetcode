package dynamicprogramming

import (
    "github.com/wujiang1994/go-leetcode/common"
    "math"
    "testing"
)

// 514. 自由之路
// https://leetcode-cn.com/problems/freedom-trail/
func Test_FindRotateSteps(t *testing.T) {
    ring := "godding"
    key := "gd"
    t.Log(findRotateSteps(ring, key))
}

// dp[i][j]表示旋转到key的第i个字符且最终落在ring的j位置的字符上的最少步骤
// 那么dp[i+1]有哪些值呢，是dp[i][j]的所有j位置能到达key的i+1位置的值
// 即dp[i+1] = min(dp[i-1][j]+h(j到i+1))(for j 存在)
func findRotateSteps(ring string, key string) int {
    length := len(ring)
    m := make(map[byte][]int)
    for i := 0; i < len(ring); i++ {
        if _, ok := m[ring[i]]; !ok {
            m[ring[i]] = make([]int, 0)
        }
        m[ring[i]] = append(m[ring[i]], i)
    }
    dp := make([][]int, len(key))
    for i := range dp {
        dp[i] = make([]int, len(ring))
    }
    for i := 0; i < len(key); i++ {
        // 初始化base
        if i == 0 {
            for _, v := range m[key[i]] {
                dp[i][v] = distance(length, 0, v) + 1
            }
            continue
        }
        for j := 0; j < len(dp[i-1]); j++ {
            if dp[i-1][j] == 0 { // 不可到达
                continue
            }
            for _, v := range m[key[i]] {
                // dp[i-1]中达到了i-1位置的j分别到达i位置的较短距离值
                if dp[i][v] == 0 {
                    dp[i][v] = dp[i-1][j] + distance(length, j, v) + 1
                } else {
                    dp[i][v] = common.Min(dp[i][v], dp[i-1][j]+distance(length, j, v)+1)
                }
            }
        }

    }
    ret := math.MaxInt32
    for _, v := range dp[len(key)-1] {
        if v != 0 && v < ret {
            ret = v
        }
    }
    return ret
}

func distance(len, start, end int) int {
    return common.Min(common.Abs(start, end), len-common.Abs(start, end))
}
