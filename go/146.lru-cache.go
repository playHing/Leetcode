/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type ListNode struct {
	Next *ListNode
	Prev *ListNode
	Key int
	Val int
}

type LRUCache struct {
	Head *ListNode
	Tail *ListNode
	Map map[int]*ListNode
	Len int
	Cap int
}


func Constructor(capacity int) LRUCache {
	m := make(map[int]*ListNode)
	cache := LRUCache{Map: m, Cap: capacity}
	return cache
}


func (this *LRUCache) Get(key int) int {
	val := -1
    if node, b := this.Map[key]; b {
		val = node.Val
		if node == this.Tail {
			return val
		}
		
		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			this.Head = node.Next
		}
		node.Next.Prev = node.Prev

		this.Tail.Next = node
		node.Prev = this.Tail
		node.Next = nil
		this.Tail = node
	}
	return val
}

func (this *LRUCache) Put(key int, value int)  {
    if this.Get(key) == -1 {
		node := &ListNode{nil, this.Tail, key, value}
		this.Map[key] = node
		if this.Tail != nil {
			this.Tail.Next = node
		} else {
			this.Head = node
		}
		this.Tail = node
		if this.Len == this.Cap {
			delete(this.Map, this.Head.Key)
			this.Head = this.Head.Next
			this.Head.Prev = nil
		} else {
			this.Len++
		}
	} else {
		this.Tail.Val = value
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

