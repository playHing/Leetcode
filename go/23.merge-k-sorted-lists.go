/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	ahead := &ListNode{0, nil}
	head := ahead
	h := []*ListNode{}
	for _, listhead := range lists {
		if listhead != nil {
			h = append(h, listhead)
		}
	}
	heap := Heap(h)
	heap.build()
	for heap.len() > 0 {
		head.Next = heap.pop()
		head = head.Next
		if head.Next != nil {
			heap.add(head.Next)
		}
	}
	return ahead.Next
}

type Heap []*ListNode

func (heap *Heap) print() {
	h := []*ListNode(*heap)
	for _, t := range h {
		fmt.Print(t.Val, " ")
	}
	fmt.Println("")
}

func (heap *Heap) len() int {
	h := []*ListNode(*heap)
	return len(h)
}

func (heap *Heap) build() {
	h := []*ListNode(*heap)
	for i := len(h)/2-1; i >= 0; i-- {
		heap.heapifyAt(i)
	}
}

func (heap *Heap) pop() *ListNode {
	h := []*ListNode(*heap)
	node := h[0]
	h[0] = h[len(h)-1]
	*heap = Heap(h[:len(h)-1])
	heap.heapifyAt(0)
	return node
}

func (heap *Heap) heapifyAt(i int) {
	h := []*ListNode(*heap)
	for {
		ni := 2*i+1
		if ni >= len(h) {
			return
		}
		if 2*i+2 < len(h) && h[2*i+2].Val < h[2*i+1].Val {
			ni = 2*i+2
		}
		if h[ni].Val < h[i].Val {
			h[i], h[ni] = h[ni], h[i]
		}
		i = ni
	}
}

func (heap *Heap) add(node *ListNode) {
	h := []*ListNode(*heap)
	h = append(h, node)
	for i := len(h)-1; i > 0; {
		pi := (i+1)/2 - 1
		if h[i].Val < h[pi].Val {
			h[i], h[pi] = h[pi], h[i]
		} else {
			break
		}
		i = pi
	}
	*heap = Heap(h)
}

// @lc code=end

