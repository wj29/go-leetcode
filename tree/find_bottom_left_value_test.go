package tree

import "testing"

// 给定一个二叉树，在树的最后一行找到最左边的值
// https://leetcode-cn.com/problems/find-bottom-left-tree-value/description/ medium
func Test_FindBottomLeftValue(t *testing.T) {

}

// 层序遍历
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0) // 构造队列
	queue = append(queue, root)   // 根结点进入队列
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:] // 出栈 最后出栈的是叶子结点的最左侧叶子结点，然后queue为空
		if root.Right != nil {
			queue = append(queue, root.Right) // 右孩子结点先入入队列，左孩子结点后进入队列，最后出来的一定是左孩子结点，
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}
