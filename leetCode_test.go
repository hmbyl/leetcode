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

func Test_swapPairs(t *testing.T) {
	listNode := newListNode([]int{1, 2, 3, 4})
	t.Log(getListNode(swapPairs(listNode)))
}

func Test_removeDuplicates(t *testing.T) {
	t.Log(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func Test_removeElement(t *testing.T) {
	t.Log(removeElement([]int{3, 2, 2, 3}, 3))
}

func Test_strStr(t *testing.T) {
	t.Log(strStr("abc", "b"))
}

func Test_pointer(t *testing.T) {
	a := make([]int, 4)
	a = append(a, 1, 2, 3)
	b := a
	t.Logf("%p %p %v %v", a, b, cap(a), cap(b))
	t.Logf("%p %p", &a, &b)
	a = append(a, 1, 2)
	t.Logf("%p %p %v %v", a, b, cap(a), cap(b))
	t.Logf("%p %p", &a, &b)
}

func Test_permute(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}

func Test_rotate(t *testing.T) {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

func Test_pow(t *testing.T) {
	t.Log(myPow(0.00001, 2147483647))
}

func Test_maxSubArray(t *testing.T) {
	t.Log(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	t.Log(maxSubArray([]int{5, 4, -1, 7, 8}))

}
