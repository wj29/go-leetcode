package tree

import (
	"github.com/wujiang1994/go-leetcode/common"
	"math"
	"testing"
)

// 124. 二叉树中的最大路径和
// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
func Test_MaxPathSum(t *testing.T) {
	root := NewTree()
	t.Log(maxPathSum(root))
}

// 找到经过该节点的最长路径长度，比较所有节点的最长路径
// 一个节点最长路径长度可以通过递归得到
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := math.MinInt32
	path4(root, &ret)
	return ret
}

// 经过该节点的最长路径和为左子树或者右子树+根节点值
func path4(root *TreeNode, val *int) int {
	if root == nil {
		return 0
	}
	left := common.Max(path4(root.Left, val), 0)
	right := common.Max(path4(root.Right, val), 0)
	*val = common.Max(*val, left+right+root.Val)
	return common.Max(left, right) + root.Val
}
