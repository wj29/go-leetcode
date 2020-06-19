package list

import "testing"

// 给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。
// 每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。
// 这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。
// 返回一个符合上述规则的链表的列表。
// 举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]
// https://leetcode-cn.com/problems/split-linked-list-in-parts
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
