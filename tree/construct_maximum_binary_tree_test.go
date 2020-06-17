package tree

import "testing"

// 给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
// 二叉树的根是数组中的最大元素。
// 左子树是通过数组中最大值左边部分构造出的最大二叉树。
// 右子树是通过数组中最大值右边部分构造出的最大二叉树。
// 通过给定的数组构建最大二叉树，并且输出这个树的根节点。
// https://leetcode-cn.com/problems/maximum-binary-tree medium
func Test_ConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	t.Log(root)
}

// 此题构建树，和根据中序前序构建二叉树，排序数组构建BST类似
// build_tree_test.go sorted_array_to_BST_test.go
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := getMaxIndex(nums)
	root := &TreeNode{
		Val: nums[index],
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func getMaxIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	maxIndex := 0
	for i, v := range nums {
		if v > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
