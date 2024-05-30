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

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcdeafb"))
}

func Test_longestPalindrome(t *testing.T) {
	//t.Log(longestPalindrome("babad"))
	//t.Log(longestPalindrome("cbbd"))
	t.Log(longestPalindrome("bb"))
}

func Test_reverse(t *testing.T) {
	t.Log(reverse(1534236469))
	t.Log(reverse(-456))
}

func Test_maxArea(t *testing.T) {
	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func Test_romanToInt(t *testing.T) {
	t.Log(romanToInt("LVIII"))
}

func Test_longestCommonPrefix(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"ab", "a"}))
}

func Test_threeSum(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func Test_letterCombinations(t *testing.T) {
	t.Log(letterCombinations("23"))
}

func Test_removeNthFromEnd(t *testing.T) {
	listNode := newListNode([]int{1, 2, 3, 4, 5})
	t.Log(getListNode(removeNthFromEnd(listNode, 5)))

}

func Test_isValid(t *testing.T) {
	t.Log(isValid("[{}]"))
	t.Log(isValid("{{]]"))
	t.Log(isValid("{()[]}"))
}

func Test_mergeTwoLists(t *testing.T) {
	l1 := newListNode([]int{1, 2, 4})
	l2 := newListNode([]int{1, 3, 4})
	t.Log(getListNode(mergeTwoLists(l1, l2)))
}

func Test_generateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
}
