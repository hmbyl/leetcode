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
最长公共前缀
https://leetcode.cn/problems/longest-common-prefix/description/
输入：strs = ["flower","flow","flight"]
输出："fl"
*/
func longestCommonPrefix(strs []string) string {
	var res string
	if len(strs) == 0 {
		return res
	}
	if len(strs) == 1 {
		return strs[0]
	}
	length := len(strs[0])
	for l := 1; l <= length; l++ {
		for i := 1; i < len(strs); i++ {
			if l > len(strs[i]) {
				return res
			}
			if strs[0][l-1] != strs[i][l-1] {
				return res
			}

		}
		res = strs[0][:l]
	}
	return res
}

/*
罗马数字转整数
https://leetcode.cn/problems/roman-to-integer/description/
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res int
	for i, _ := range s {
		letter := s[i]
		num := m[letter]
		if i > 0 {
			lastNum := m[s[i-1]]
			if num > lastNum {
				res -= 2 * lastNum
			}
			res += num
		} else {
			res += num
		}
	}
	return res
}

/*
盛最多水的容器
https://leetcode.cn/problems/container-with-most-water/description/
输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	m := 0
	for left < right {
		h1 := height[left]
		h2 := height[right]
		m = max(m, min(h1, h2)*(right-left))
		if h1 < h2 {
			left++
		} else {
			right--
		}
	}
	return m
}

/*
整数反转
https://leetcode.cn/problems/reverse-integer/description/
输入：x = 123
输出：321
*/
func reverse(x int) int {
	var res int
	var flag bool
	if x < 0 {
		flag = true
		x = -x
	}
	for x > 0 {
		tmp := x % 10
		res = res*10 + tmp
		x = x / 10
	}
	if flag {
		res = -res
	}
	if res < -1<<31 || res > (1<<31-1) {
		res = 0
	}
	return res
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
