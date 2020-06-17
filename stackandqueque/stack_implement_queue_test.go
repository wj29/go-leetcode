package stackandqueque

import (
	"fmt"
	"testing"
)

// 使用栈实现队列的下列操作:
// push(x) -- 将一个元素放入队列的尾部
// pop() -- 从队列首部移除元素
// peek() -- 返回队列首部的元素
// empty() -- 返回队列是否为空
// https://leetcode-cn.com/problems/implement-queue-using-stacks/description/ easy
func Test_StackImplementQueue(t *testing.T) {
	obj := ConstructorQueue()
	obj.Push(1)
	fmt.Println(obj)
	obj.Push(2)
	obj.Push(3)
	param2 := obj.Peek()
	t.Log(param2)
	param1 := obj.Pop()
	t.Log(param1)
	param3 := obj.Empty()
	t.Log(param3)
}

// go既没有栈也没有队列
type MyQueue struct {
	Elem  []int
}

/** Initialize your data structure here. */
func ConstructorQueue() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Elem = append(this.Elem, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		panic("empty queue")
	}
	val := this.Elem[0]
	this.Elem = this.Elem[1:]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		panic("empty queue")
	}
	return this.Elem[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Elem) == 0
}
