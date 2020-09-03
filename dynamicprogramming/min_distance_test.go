package dynamicprogramming

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 72. 编辑距离
// https://leetcode-cn.com/problems/edit-distance/
func Test_MinDistance(t *testing.T) {
	word1 := "intention"
	word2 := "execution"
	t.Log(minDistance(word1, word2))
}

// dp[i][j]表示从word1的0-i-1位变成word2的0-j-1位需要的操作数
// 由于有三种操作
// - 插入,对应操作dp[i-1][j]到dp[i][j]
// - 删除,dp[i][j-1]到dp[i][j]
// - 修改,dp[i-1][j-1]到dp[i][j]
// 最小操作步数取上面最小的操作结果
// base dp[0][j]需要j步 dp[i][0]需要i
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		for j := 0; j < len(word2)+1; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				modifyCost := 0
				if word1[i-1] != word2[j-1] {
					modifyCost++
				}
				dp[i][j] = common.Min(dp[i-1][j-1]+modifyCost, common.Min(dp[i-1][j], dp[i][j-1])+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
