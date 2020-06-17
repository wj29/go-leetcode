package stackandqueque

import "testing"

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈
// push(x) —— 将元素 x 推入栈中
// pop() —— 删除栈顶的元素
// top() —— 获取栈顶元素
// getMin() —— 检索栈中的最小元素
// https://leetcode-cn.com/problems/min-stack  easy
func Test_MinStack(t *testing.T) {
	minStack := Constructor1()
	minStack.Push(2)
	minStack.Push(0)
	minStack.GetMin()
	minStack.Pop()
}

type MinStack struct {
	Elem []int
	min  int
}

func Constructor1() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.Elem) == 0 {
		this.min = x
	} else {
		this.min = min(this.min, x)
	}
	this.Elem = append(this.Elem, x)
}

func (this *MinStack) Pop() {
	this.Elem = this.Elem[:len(this.Elem)-1]
	this.min = getMin(this.Elem)
}

func (this *MinStack) Top() int {
	return this.Elem[len(this.Elem)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
