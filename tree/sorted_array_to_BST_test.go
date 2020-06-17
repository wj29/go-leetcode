package tree

import (
	"testing"
)

// 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1
// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/  easy
func Test_SortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	t.Log(root)
}

// 二分得到
// 二叉搜索树的中序遍历即是升序的有序数组 [ [左子树遍历结果] 根 [右子树遍历结果] ]
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	l, r := 0, len(nums)-1
	mid := l + (r-l)>>1
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
