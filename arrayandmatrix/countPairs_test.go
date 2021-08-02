package arrayandmatrix

import (
    "math"
    "testing"
)

func Test_CountPairs(t *testing.T) {
    deliciousness := []int{1, 3, 5, 7, 9}
    deliciousness = []int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}
    t.Log(countPairs(deliciousness))
}

func countPairs(deliciousness []int) (ans int) {
    maxVal := deliciousness[0]
    for _, val := range deliciousness[1:] {
        if val > maxVal {
            maxVal = val
        }
    }
    maxSum := maxVal * 2
    cnt := map[int]int{}
    for _, val := range deliciousness {
        for sum := 1; sum <= maxSum; sum <<= 1 { // 以2的次方数为循环条件，减少循环
            ans += cnt[sum-val] // 用排列组合求种类数
        }
        cnt[val]++
    }
    return ans % (1e9 + 7)
}

// 暴力解法超时
func countPairs2(deliciousness []int) int {
    count := 0
    deliciousMap := make(map[int]bool)
    for i := 0; i < len(deliciousness); i++ {
    loop:
        for j := i + 1; j < len(deliciousness); j++ {
            if deliciousMap[deliciousness[i]+deliciousness[j]] {
                count++
                continue loop
            }
            for k := 0; k <= 21; k++ {
                if (deliciousness[i]+deliciousness[j])^1<<k == 0 {
                    deliciousMap[deliciousness[i]+deliciousness[j]] = true
                    count++
                    continue loop
                }
            }
            deliciousMap[deliciousness[i]+deliciousness[j]] = false
        }
    }
    return count % (int(math.Pow(10, 9)) + 7)
}
