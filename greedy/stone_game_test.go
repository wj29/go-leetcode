package greedy

import "testing"

// A和B用几堆石子做游戏，偶数堆石子，每堆有正证书个石子plies[i]
// 游戏胜负以石子个数判断，石子个数是奇数，不会是平局
// A和B轮流进行，每次从行首或者行尾取走石子，AB都希望自己获胜，A胜返回true，B胜返回false
// https://leetcode-cn.com/problems/stone-game/ medium
// eg: 输入 [5,3,4,5] 输出 true
// TODO

func Test_StoneGame(t *testing.T) {

}

// 获胜的条件 = 最后一次取出的石子个数+前n-1次取出石子个数的总和
func stoneGame(piles []int) bool {
	return false
}
