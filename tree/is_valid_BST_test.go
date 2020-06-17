package tree

import (
	"fmt"
	"math"
	"testing"
)

// 判断一棵树是不是二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/  medium
func Test_IsValidBST(t *testing.T) {
	//root1 := NewBSTree()
	//t.Log(isValidBST1(root1))

	root2 := NewWrongBSTree()
	t.Log(isValidBST1(root2))
}

// 迭代
// 二叉搜索树中序遍历是有序的，通过遍历后的数组判断即可
// 在中序遍历的过程中直接判断
func isValidBST(root *TreeNode) bool {
	var (
		res   []int
		stack []*TreeNode
	)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1 // 栈顶
		res = append(res, stack[index].Val)
		if len(res) > 1 && res[len(res)-1] < res[len(res)-2] {
			return false
		}
		root = stack[index].Right
		stack = stack[:index]
	}
	return true
}

// 递归
// 一个二叉搜索树具有如下特征：
// 节点的左子树只包含小于当前节点的数
// 节点的右子树只包含大于当前节点的数
// 所有左子树和右子树自身必须也是二叉搜索树
func isValidBST1(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		fmt.Println("root:", root, "low:", lower, "up:", upper)
		return false
	}
	fmt.Println("root:", root, "left:", root.Left, "right:", root.Right, "low:", lower, "up:", upper)
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
