package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListNode(l *ListNode) []int {
	var nums []int
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}

func newListNode(nums []int) *ListNode {
	head := new(ListNode)
	node := head
	for i, v := range nums {
		node.Val = v
		if i != len(nums)-1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	return head
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
