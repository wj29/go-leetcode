package tree

// 前序遍历的代码在进入某一个节点之前的那个时间点执行，后序遍历代码在离开某个节点之后的那个时间点执行
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func NewTree() *TreeNode {
	p1 := &TreeNode{
		Val: 1,
	}
	p2 := &TreeNode{
		Val: 2,
	}
	p3 := &TreeNode{
		Val: 3,
	}
	p4 := &TreeNode{
		Val: 4,
	}
	p5 := &TreeNode{
		Val: 5,
	}
	p6 := &TreeNode{
		Val: 6,
	}
	p1.Left = p2
	p1.Right = p3
	p2.Left = p4
	p2.Right = p5
	p3.Left = p6
	return p1
}

func NewSmallBSTree() *TreeNode {
	p1 := &TreeNode{
		Val: 5,
	}
	p2 := &TreeNode{
		Val: 2,
	}
	p3 := &TreeNode{
		Val: 13,
	}
	p1.Left = p2
	p1.Right = p3
	return p1
}

func NewBSTree() *TreeNode {
	p1 := &TreeNode{
		Val: 4,
	}
	p2 := &TreeNode{
		Val: 2,
	}
	p3 := &TreeNode{
		Val: 5,
	}
	p4 := &TreeNode{
		Val: 1,
	}
	p5 := &TreeNode{
		Val: 7,
	}
	p6 := &TreeNode{
		Val: 3,
	}
	p1.Left = p2
	p1.Right = p3
	p2.Left = p4
	p2.Right = p6
	p3.Right = p5
	return p1
}

func NewWrongBSTree() *TreeNode {
	p1 := &TreeNode{
		Val: 4,
	}
	p2 := &TreeNode{
		Val: 3,
	}
	p3 := &TreeNode{
		Val: 5,
	}
	p4 := &TreeNode{
		Val: 1,
	}
	p5 := &TreeNode{
		Val: 7,
	}
	p6 := &TreeNode{
		Val: 6,
	}
	p1.Left = p2
	p1.Right = p3
	p2.Left = p4
	p2.Right = p6
	p3.Right = p5
	return p1
}
