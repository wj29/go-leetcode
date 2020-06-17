package leetcode

import (
	"testing"
)

// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
// https://leetcode-cn.com/problems/lru-cache  medium
// 进阶: 时间复杂度为1
func Test_LRUCache(t *testing.T) {

}

// 通过双向链表实现插入更新时间复杂度1
// map实现查找时间复杂度1
type LRUCache struct { // 双向链表
	cap  int
	head *lruNode
	tail *lruNode
	m    map[int]*lruNode // map
}

type lruNode struct {
	key        int
	value      int
	prev, next *lruNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap: capacity,
		head: &lruNode{},
		tail: &lruNode{},
		m:   make(map[int]*lruNode),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; ok {
		// 访问到的结点移动到表头
		this.remove(this.m[key]) // 删除该结点
		this.setFront(this.m[key])
		return this.m[key].value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 { // 存在
		this.m[key].value = value  // 改值
		this.remove(this.m[key])   // 删除该结点
		this.setFront(this.m[key]) // 移动到表头
	} else { // 不存在
		this.m[key] = &lruNode{
			key:   key,
			value: value,
		}
		this.setFront(this.m[key])
		if len(this.m) > this.cap { // lru缓存满了，删除一个
			lastNode := this.tail.prev
			this.remove(lastNode)
			delete(this.m, lastNode.key)
		}
	}
}

func (this *LRUCache) remove(node *lruNode) {
	// 移除该结点
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) setFront(node *lruNode) {
	// 放到到表头
	tmp := this.head.next // 记录原表头
	this.head.next = node // 将head指向新表头
	node.next = tmp       // 连接新表头和原表头
	tmp.prev = node       // 原表头前指针指向新表头
	node.prev = this.head // 将新表头指向head
}
