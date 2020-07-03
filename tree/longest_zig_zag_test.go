package tree

import (
	"github.com/wujiang1994/go-leetcode/common"
	"math"
	"testing"
)

// 给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：
// 选择二叉树中 任意 节点和一个方向（左或者右）。
// 如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
// 改变前进方向：左变右或者右变左。
// 重复第二步和第三步，直到你在树中无法继续移动。
// 交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。
// 请你返回给定树中最长 交错路径 的长度。
// https://leetcode-cn.com/problems/longest-zigzag-path-in-a-binary-tree
func Test_LongestZigZag(t *testing.T) {
	root := NewTree()
	t.Log(longestZigZag(root))
	t.Log(longestZigZag2(root))
}

// 递归，遍历所有的节点
// 一个树的最长交错路径=max(左子树，右子树)的交错路径+1
// 超时
func longestZigZag(root *TreeNode) int {
	ret := 0
	if root == nil || (root.Left == nil && root.Right == nil) {
		return ret
	}
	traversalForZigZag(root, &ret)
	return ret
}

func traversalForZigZag(root *TreeNode, ret *int) {
	if root == nil {
		return
	}
	val := common.Max(zigZag(root.Left, true), zigZag(root.Right, false))
	if val > *ret {
		*ret = val
	}
	traversalForZigZag(root.Left, ret)
	traversalForZigZag(root.Right, ret)
}

func zigZag(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	if isLeft {
		max = common.Max(zigZag(root.Right, false), max)
	} else {
		max = common.Max(zigZag(root.Left, true), max)
	}
	return max + 1
}

// 上述的解法超时在于，遍历了每一个节点，然后遍历每一个节点所有子节点，重复遍历了多次
// 此解法中对于一个节点，它可以是上一个节点的左或者右子节点，假如它是上一个节点的左子节点，那么它的值应该是其右子树值+1
// 同理当它是右子节点，它的值是其左子树值+1
// 即每个节点都有两个交错值，分别是左和右
// root可以是上一个节点(虚构)的左子节点或右子节点
func longestZigZag2(root *TreeNode) int {
	ret := 0
	if root == nil || (root.Left == nil && root.Right == nil) {
		return ret
	}
	traversalForZigZag2(root, &ret, true)
	traversalForZigZag2(root, &ret, false)
	return ret
}

func traversalForZigZag2(root *TreeNode, ret *int, isLeft bool) int {
	if root == nil {
		return 0
	}
	right := traversalForZigZag2(root.Right, ret, false)
	left := traversalForZigZag2(root.Left, ret, true)
	*ret = common.Max(common.Max(left, right), *ret)
	if isLeft {
		return right + 1
	}
	return left + 1
}
