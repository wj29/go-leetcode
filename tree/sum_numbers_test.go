package tree

import (
	"strconv"
	"testing"
)

// 129. 求根到叶子节点数字之和
// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
func Test_SumNumbers(t *testing.T) {
	root := NewSmallBSTree()
	t.Log(sumNumbers(root))
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := make([]string, 0)
	traversalForSumNumbers(root, &ret, strconv.Itoa(root.Val))
	r := 0
	for _, v := range ret {
		tmp, _ := strconv.Atoi(v)
		r += tmp
	}
	return r
}

func traversalForSumNumbers(root *TreeNode, ret *[]string, prev string) {
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, prev)
		return
	}
	if root.Left != nil {
		traversalForSumNumbers(root.Left, ret, prev+strconv.Itoa(root.Left.Val))
	}
	if root.Right != nil {
		traversalForSumNumbers(root.Right, ret, prev+strconv.Itoa(root.Right.Val))
	}
}
