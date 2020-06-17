package list

import "testing"

func Test_SplitListToParts(t *testing.T) {
	root := NewList()
	ret := splitListToParts(root, 4)
	for root != nil {
		t.Log(root.Val)
		root = root.Next
	}
	for _, v := range ret {
		t.Log("[")
		for v != nil {
			t.Log(v.Val)
			v = v.Next
		}
		t.Log("]")
	}
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	length := 0
	current := root
	for current != nil {
		current = current.Next
		length++
	}
	nums, count := countAndNums(length, k)
	ret := make([]*ListNode, 0)
	current = root
	p := 0
	sum := 0
	flag := true
	for current != nil {
		if len(ret) == count && flag {
			nums -= 1
			flag = false
		}
		if p == sum {
			ret = append(ret, current)
			sum += nums
		}
		if p == sum-1 {
			newCurrent := current.Next
			current.Next = nil
			current = newCurrent
		} else {
			current = current.Next
		}
		p++
	}
	for k-len(ret) > 0 {
		ret = append(ret, nil)
	}
	return ret
}

func countAndNums(len, k int) (int, int) {
	nums := len/k + 1 // 每个子数组多一个的有多少结点
	count := len % k  // 子结点数多一个的应有多少个
	return nums, count
}
