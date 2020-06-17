package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList() *ListNode {
	p1 := &ListNode{
		Val: 1,
	}
	p2 := &ListNode{
		Val: 2,
	}
	p3 := &ListNode{
		Val: 3,
	}
	p4 := &ListNode{
		Val: 4,
	}
	p5 := &ListNode{
		Val: 5,
	}
	p1.Next = p2
	p2.Next = p3
	p3.Next = p4
	p4.Next = p5
	return p1
}

func NewCycleList() *ListNode {
	p1 := &ListNode{
		Val: 1,
	}
	p2 := &ListNode{
		Val: 2,
	}
	p3 := &ListNode{
		Val: 3,
	}
	p4 := &ListNode{
		Val: 4,
	}
	p5 := &ListNode{
		Val: 5,
	}
	p1.Next = p2
	p2.Next = p3
	p3.Next = p4
	p4.Next = p5
	p5.Next = p2
	return p1
}
