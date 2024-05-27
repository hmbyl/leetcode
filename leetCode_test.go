package main

import "testing"

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	t.Log(twoSum(nums, target))
	nums = []int{3, 2, 4}
	target = 6
	t.Log(twoSum(nums, target))
}

func Test_addTwoNumbers(t *testing.T) {
	l1 := newListNode([]int{2, 4, 3})
	l2 := newListNode([]int{5, 6, 4})
	t.Log(getListNode(addTwoNumbers(l1, l2)))
	l1 = newListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = newListNode([]int{9, 9, 9, 9})
	t.Log(getListNode(addTwoNumbers(l1, l2)))

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

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcdeafb"))
}