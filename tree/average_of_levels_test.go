package tree

import (
	"testing"
)

// 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组
// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/description/  easy
func Test_AverageOfLevels(t *testing.T) {
	root := NewTree()
	t.Log(averageOfLevels(root))
}

type mapValue struct {
	*TreeNode
	mark int
}

// 层序遍历
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := make([]*mapValue, 0) // 用于层序遍历
	queue = append(queue, &mapValue{
		TreeNode: root,
		mark:     0,
	})
	arr := make([][]*TreeNode, 0)        // 下标代表第几层，数组代表该层结点集合
	arr = append(arr, []*TreeNode{root}) // 0下标只有root
	for i := 0; i < len(queue); i++ {
		root := queue[i]
		if root.Left != nil {
			queue = append(queue, &mapValue{
				TreeNode: root.Left,
				mark:     root.mark + 1,
			})
		}
		if root.Right != nil {
			queue = append(queue, &mapValue{
				TreeNode: root.Right,
				mark:     root.mark + 1,
			})
		}
	}
	ret := make([]float64, queue[len(queue)-1].mark+1)
	count := make([]int, queue[len(queue)-1].mark+1)
	for i := 0; i < len(queue); i++ {
		ret[queue[i].mark] += float64(queue[i].Val)
		count[queue[i].mark]++
	}
	for i := range ret {
		ret[i] /= float64(count[i])
	}
	return ret
}
