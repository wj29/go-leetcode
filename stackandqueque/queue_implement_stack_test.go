package stackandqueque

import (
	"fmt"
	"testing"
)

// 使用队列实现栈的下列操作:
// push(x) -- 元素 x 入栈
// pop() -- 移除栈顶元素
// top() -- 获取栈顶元素
// empty() -- 返回栈是否为空
// https://leetcode-cn.com/problems/implement-stack-using-queues  easy
func Test_QueueImplementStack(t *testing.T) {
	obj := ConstructorStack()
	obj.Push(1)
	fmt.Println(obj)
	obj.Push(2)
	obj.Push(3)
	param1 := obj.Pop()
	t.Log(param1)
	param3 := obj.Empty()
	t.Log(param3)
}

type MyStack struct {
	Elem []int
}

/** Initialize your data structure here. */
func ConstructorStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Elem = append(this.Elem, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.Elem[len(this.Elem)-1]
	this.Elem = this.Elem[:len(this.Elem)-1]
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Elem[len(this.Elem)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Elem) == 0
}
