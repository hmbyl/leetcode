package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

/*
两数之和
https://leetcode.cn/problems/two-sum/submissions/217515930/
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for j, num := range nums {
		if i, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = j
	}
	return []int{}
}

/*
两数相加
https://leetcode.cn/problems/add-two-numbers/description/
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	node := head
	var carry int
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		sum = sum - carry*10
		node.Val = sum
		if l1 != nil || l2 != nil {
			node.Next = new(ListNode)
			node = node.Next
		}

	}
	if carry > 0 {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = carry
	}
	return head
}

/*
无重复字符的最长子串
https://leetcode.cn/problems/longest-substring-without-repeating-characters/
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var res int
	rk := 1
	for i, _ := range s {
		m[s[i]] = 1
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk < len(s) && m[s[rk]] == 0 {
			m[s[rk]] = 1
			rk++
		}
		res = max(res, rk-i)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

/*
最长回文子串
https://leetcode.cn/problems/longest-palindromic-substring/description/
示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	begin := 0
	length := 1
	for strLength := 1; strLength <= n; strLength++ {

		for i, _ := range s {
			j := i + strLength - 1
			if j >= n {
				break
			}
			if strLength <= 3 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else if dp[i+1][j-1] == 1 && s[i] == s[j] {
				dp[i][j] = 1
			}
			if dp[i][j] == 1 && strLength > length {
				length = strLength
				begin = i
			}
		}
	}
	return s[begin : begin+length]
}
