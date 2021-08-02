package stackandqueque

import (
    "testing"
)

// 85. 最大矩形
// https://leetcode-cn.com/problems/maximal-rectangle/
func Test_MaximalRectangle(t *testing.T) {
    matrix := [][]byte{
        {'1', '0', '1', '0', '0'},
        {'1', '0', '1', '1', '1'},
        {'1', '1', '1', '1', '1'},
        {'1', '0', '0', '1', '0'},
    }
    t.Log(maximalRectangle(matrix))
}

// 求最大矩形
// 本题和84题本质上解法是一致的，将每一行看成一个柱状图，套用84的解法
func maximalRectangle(matrix [][]byte) int {
    return 0
}
