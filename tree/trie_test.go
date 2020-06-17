package tree

import "testing"

// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/ medium
func Test_Trie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	t.Log(trie.Search("apple"))   // 返回 true
	t.Log(trie.Search("app"))     // 返回 false
	t.Log(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	t.Log(trie.Search("app")) // 返回 true
}

type Trie struct {
	children map[string]*Trie // 孩子结点
	val      string           // 当前结点值
	flag     bool             // 标志当前结点是不是字符串
}

func Constructor() Trie {
	trie := Trie{
		children: make(map[string]*Trie),
		val:      "",
	}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for i := range word {
		if _, ok := this.children[string(word[i])]; !ok {
			this.children[string(word[i])] = &Trie{
				children: make(map[string]*Trie),
				val:      string(word[i]),
			}
		}
		this = this.children[string(word[i])]
		if i == len(word)-1 {
			this.flag = true
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := range word {
		if _, ok := this.children[string(word[i])]; !ok {
			return false
		}
		this = this.children[string(word[i])]
	}
	return this.flag
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		if _, ok := this.children[string(prefix[i])]; !ok {
			return false
		}
		this = this.children[string(prefix[i])]
	}
	return true
}
