package tree

import "testing"

// 计算给定二叉树的所有左叶子之和
// https://leetcode-cn.com/problems/sum-of-left-leaves/description/  easy
func Test_SumOfLeftLeaves(t *testing.T) {
	root := NewTree()
	t.Log(sumOfLeftLeaves(root))
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	traversal1(root, &sum)
	return sum
}

// 判断该节点是否为左叶子结点，通过全局变量sum记录总和
func traversal1(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Left.Val
	}
	traversal1(root.Left, sum)
	traversal1(root.Right, sum)
}
