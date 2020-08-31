package tree

import (
	"fmt"
	"testing"
)

// 102. 二叉树的层序遍历
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 107. 二叉树的层次遍历 II
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 429. N叉树的层序遍历
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

func Test_LevelOrder(t *testing.T) {
	root := NewTree()
	t.Log(levelOrder(root))

	a := []int{1, 2}
	b := []int{3, 5, 4, 1, 1, 1}
	c := []int{8, 8, 8, 8, 8}
	a = b
	b = c
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	pre := []*TreeNode{root}
	for i := 0; len(pre) > 0; i++ {
		ret = append(ret, []int{})
		tmp := make([]*TreeNode, 0)
		for j := 0; j < len(pre); j++ {
			ret[i] = append(ret[i], pre[j].Val)
			if pre[j].Left != nil {
				tmp = append(tmp, pre[j].Left)
			}
			if pre[j].Right != nil {
				tmp = append(tmp, pre[j].Right)
			}
		}
		pre = tmp
	}
	return ret
}
