package tree

import "testing"

// 116. 填充每个节点的下一个右侧节点指针
// 完美二叉树
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
func Test_Connect(t *testing.T) {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

// 层序遍历
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	pre := []*Node{root}
	for i := 0; len(pre) > 0; i++ {
		tmp := make([]*Node, 0)
		for j := 0; j < len(pre)-1; j++ {
			pre[j].Next = pre[j+1]
			if pre[j].Left != nil {
				tmp = append(tmp, pre[j].Left)
			}
			if pre[j].Right != nil {
				tmp = append(tmp, pre[j].Right)
			}
		}
		if pre[len(pre)-1].Left != nil {
			tmp = append(tmp, pre[len(pre)-1].Left)
		}
		if pre[len(pre)-1].Right != nil {
			tmp = append(tmp, pre[len(pre)-1].Right)
		}
		pre = tmp
	}
	return root
}
