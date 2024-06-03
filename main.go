package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

/*
二叉树展开为链表
https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
*/
func flatten(root *TreeNode) {
	cur := root
	var res []int
	var backTrack func(node *TreeNode)
	backTrack = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		backTrack(node.Left)
		backTrack(node.Right)
	}
	backTrack(root)
	for i, v := range res {
		cur.Val = v
		if i != len(res)-1 {
			cur.Right = new(TreeNode)
			cur.Left = nil
			cur = cur.Right
		}
	}
}

/*
二叉树路径总和 二
https://leetcode.cn/problems/path-sum-ii/
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	var backTrack func(node *TreeNode, tmpSum int, tmpNodes []int)

	var res [][]int
	backTrack = func(node *TreeNode, tmpSum int, tmpNodes []int) {
		if node == nil {
			return
		}
		val := node.Val
		tmpSum = tmpSum + val
		tmpNodes = append(tmpNodes, val)
		if tmpSum == targetSum && node.Left == nil && node.Right == nil {
			res = append(res, tmpNodes)
			return
		}
		newTmpNodes1 := make([]int, len(tmpNodes))
		copy(newTmpNodes1, tmpNodes)
		newTmpNodes2 := make([]int, len(tmpNodes))
		copy(newTmpNodes2, tmpNodes)
		backTrack(node.Left, tmpSum, newTmpNodes1)
		backTrack(node.Right, tmpSum, newTmpNodes2)
	}
	backTrack(root, 0, []int{})
	return res
}

/*
二叉树路径总和
https://leetcode.cn/problems/path-sum/
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	var backTrack func(node *TreeNode, tmpSum int)

	var res bool
	backTrack = func(node *TreeNode, tmpSum int) {
		if node == nil {
			return
		}
		val := node.Val
		tmpSum = tmpSum + val
		if tmpSum == targetSum && node.Left == nil && node.Right == nil {
			res = true
			return
		}
		backTrack(node.Left, tmpSum)
		backTrack(node.Right, tmpSum)
	}
	backTrack(root, 0)
	return res
}

/*
二叉树最小深度
https://leetcode.cn/problems/minimum-depth-of-binary-tree/
*/
func minDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		depth++
		var newQueue []*TreeNode
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}
	return depth
}

/*
将有序数组转化为二叉搜索树
https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var backTrack func(left, right int) *TreeNode
	backTrack = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := right - (right-left)/2
		node := &TreeNode{Val: nums[mid]}
		node.Left = backTrack(left, mid-1)
		node.Right = backTrack(mid+1, right)
		return node
	}
	return backTrack(0, len(nums)-1)
}

/*
二叉树的层序遍历，自底向上
https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
*/
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var newQueue []*TreeNode
		var tmpRes []int
		for _, node := range queue {
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		res = append(res, tmpRes)
		queue = newQueue
	}
	var ret [][]int
	for i := len(res) - 1; i >= 0; i-- {
		ret = append(ret, res[i])
	}
	return ret
}

/*
二叉树的最大深度
*/
func maxDepth(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var newQueue []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		res++
		queue = newQueue
	}
	return res
}

/*
二叉树的程序遍历
https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
*/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var newQueue []*TreeNode
		var tmpRes []int
		for _, node := range queue {
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		res = append(res, tmpRes)
		queue = newQueue
	}
	return res
}

/*
对称二叉树
https://leetcode.cn/problems/symmetric-tree/
*/
func isSymmetric(root *TreeNode) bool {
	var backTrack func(left, right *TreeNode) bool
	backTrack = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return backTrack(left.Left, right.Right) && backTrack(left.Right, right.Left)
	}
	return backTrack(root, root)
}

/*
验证二叉搜索树
https://leetcode.cn/problems/validate-binary-search-tree/description/
输入：root = [2,1,3]
输出：true
*/
func isValidBST(root *TreeNode) bool {
	var backTrack func(node *TreeNode, lower, upper int64) bool
	backTrack = func(node *TreeNode, lower, upper int64) bool {
		if node == nil {
			return true
		}
		val := int64(node.Val)
		if val >= upper || val <= lower {
			return false
		}
		return backTrack(node.Left, lower, int64(node.Val)) && backTrack(node.Right, int64(node.Val), upper)
	}
	return backTrack(root, math.MinInt64, math.MaxInt64)

}

/*
二叉树的中序遍历
*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var backTrack func(node *TreeNode)
	backTrack = func(node *TreeNode) {
		if node == nil {
			return
		}
		backTrack(node.Left)
		res = append(res, node.Val)
		backTrack(node.Right)
	}
	backTrack(root)
	return res
}

/*
子集
https://leetcode.cn/problems/subsets-ii/description/
示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
*/
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	res = append(res, []int{})
	if len(nums) == 0 {
		return res
	}
	m := make(map[string]int)
	var backTrack func(length, beginIdx int, tmp []int)
	backTrack = func(length, beginIdx int, tmp []int) {
		if length == len(tmp) {
			if beginIdx > len(nums) {
				return
			}
			sort.Ints(tmp)
			var key string
			for _, v := range tmp {
				key += fmt.Sprintf("%v", v)
			}
			if _, ok := m[key]; !ok {
				res = append(res, tmp)
				m[key] = 1
			}
			return
		}
		for i := beginIdx; i < len(nums); i++ {
			newTmp := make([]int, len(tmp))
			copy(newTmp, tmp)
			newTmp = append(newTmp, nums[i])
			backTrack(length, i+1, newTmp)
		}

	}
	for length := 1; length <= l; length++ {
		backTrack(length, 0, []int{})
	}
	//backTrack(4, 0, []int{})
	return res
}

/*
合并两个有序数组
https://leetcode.cn/problems/merge-sorted-array/description/
示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

/*
矩阵置零
https://leetcode.cn/problems/set-matrix-zeroes/description/
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
*/
func setZeroes(matrix [][]int) {
	var xArr, yArr []int
	for y, row := range matrix {
		for x, num := range row {
			if num == 0 {
				xArr = append(xArr, x)
				yArr = append(yArr, y)
			}
		}
	}
	for y, row := range matrix {
		for x, _ := range row {
			for _, tmpX := range xArr {
				if x == tmpX {
					matrix[y][x] = 0
				}
			}
			for _, tmpY := range yArr {
				if y == tmpY {
					matrix[y][x] = 0
				}
			}
		}
	}
}

/*
爬楼梯
https://leetcode.cn/problems/climbing-stairs/description/
示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if n >= 2 {
		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}

/*
二进制求和
https://leetcode.cn/problems/add-binary/description/
示例 1：

输入:a = "11", b = "1"
输出："100"
示例 2：

输入：a = "1010", b = "1011"
输出："10101"
*/

func addBinary(a string, b string) string {
	idxA := len(a) - 1
	idxB := len(b) - 1
	carry := 0
	var res string
	for idxA >= 0 || idxB >= 0 {
		aNum := 0
		bNum := 0
		if idxA >= 0 {
			aNum = int(a[idxA] - '0')
			idxA--
		}
		if idxB >= 0 {
			bNum = int(b[idxB] - '0')
			idxB--
		}
		sum := aNum + bNum + carry
		res = fmt.Sprintf("%v%v", sum%2, res)
		carry = sum / 2
	}
	if carry > 0 {
		res = fmt.Sprintf("%v%v", carry, res)

	}
	return res
}

/*
加1
示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]

https://leetcode.cn/problems/plus-one/description/
*/
func plusOne(digits []int) []int {
	carry := 0
	for idx := len(digits) - 1; idx >= 0; idx-- {
		num := digits[idx]
		sum := num + carry
		if idx == len(digits)-1 {
			sum++
		}
		carry = sum / 10
		digits[idx] = sum % 10
	}
	var res []int
	if carry > 0 {
		res = make([]int, len(digits)+1)
		res[0] = carry
		copy(res[1:], digits)
	} else {
		res = make([]int, len(digits))
		copy(res, digits)
	}
	return res
}

/*
跳跃游戏
https://leetcode.cn/problems/jump-game/solutions/203549/tiao-yue-you-xi-by-leetcode-solution/
示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
*/
func canJump(nums []int) bool {
	rightMost := 0
	for i, _ := range nums {
		if i > 0 {
			rightMost = max(rightMost, i-1+nums[i-1])
			if rightMost < i {
				return false
			}
		}
	}
	return true
}

/*
螺旋矩阵
https://leetcode.cn/problems/spiral-matrix/description/
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

*/
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	maxY := len(matrix) - 1
	maxX := len(matrix[0]) - 1
	minY, minX := 0, 0
	x, y := -1, 0
	direct := ""
	total := (maxX + 1) * (maxY + 1)
	for len(res) < total {

		if direct == "" {
			direct = "right"
			minY++
		} else if direct == "right" && x == maxX {
			maxX--
			direct = "down"
		} else if direct == "down" && y == maxY {
			maxY--
			direct = "left"
		} else if direct == "left" && x == minX {
			minX++
			direct = "up"
		} else if direct == "up" && y == minY {
			minY++
			direct = "right"
		}
		switch direct {
		case "right":
			x++
		case "down":
			y++
		case "left":
			x--
		case "up":
			y--
		}
		val := matrix[y][x]
		res = append(res, val)

	}
	return res
}

/*
最大子数组和
https://leetcode.cn/problems/maximum-subarray/
示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6
*/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	m := nums[0]
	dp[0] = m
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		m = max(m, dp[i])
	}
	return m
}

/*
https://leetcode.cn/problems/powx-n/description/
输入：x = 2.00000, n = 10
输出：1024.00000
输入：x = 2.00000, n = -2
输出：0.25000
*/
func myPow(x float64, n int) float64 {
	flag := false
	if n < 0 {
		n = -n
		flag = true
	}
	var res float64
	if n == 1 {
		res = x
	} else if n == 0 {
		return 1
	} else {
		tmpRes := myPow(x, n/2)
		if n%2 == 0 {
			res = tmpRes * tmpRes
		} else {
			res = tmpRes * tmpRes * x
		}
	}
	if flag {
		return 1 / res
	}
	return res
}

/*
旋转图像
https://leetcode.cn/problems/rotate-image/
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]
*/
func rotate(matrix [][]int) {
	l := len(matrix)
	newMtrix := make([][]int, l)
	for i, _ := range newMtrix {
		newMtrix[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := l - 1; j >= 0; j-- {
			v := matrix[j][i]
			newMtrix[i][l-1-j] = v
		}
	}
	fmt.Println(newMtrix)
	copy(matrix, newMtrix)
}

/*
全排列
https://leetcode.cn/problems/permutations/description/
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permute(nums []int) [][]int {
	var res [][]int
	var backTrack func(ignoreIndex, tmp []int)
	backTrack = func(ignoreIndexArr, tmp []int) {
		if len(tmp) == len(nums) {
			res = append(res, tmp)
			return
		}
		for i, v := range nums {
			find := false
			for _, ignoreIndex := range ignoreIndexArr {
				if ignoreIndex == i {
					find = true
				}
			}
			if !find {
				newIgnoreIndexArr := append(ignoreIndexArr, i)
				newTmp := append(tmp, v)
				backTrack(newIgnoreIndexArr, newTmp)
			}
		}
	}
	backTrack([]int{}, []int{})
	return res
}

/*
字符串相乘
https://leetcode.cn/problems/multiply-strings/description/
输入: num1 = "123", num2 = "12"
输出: "1476" 123*10=1230+123*2=246 =1476
*/
func multiply(num1 string, num2 string) string {
	return ""
}

/*
找出字符串中第一个匹配项的下标
https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0
*/
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if haystack == needle {
		return 0
	}
	idx := -1
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for q, v2 := range needle {
			v1 := haystack[q+i]
			if v1 != byte(v2) {
				idx = -1
				break
			} else if idx == -1 {
				idx = i
			}
			if q == len(needle)-1 {
				return idx
			}
		}
	}
	return idx
}

/*
移除元素
https://leetcode.cn/problems/remove-element/description/
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2,_,_]
*/
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

/*
删除有序数组中的重复项
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	slow := 0
	fast := 1
	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			nums[slow+1] = nums[fast]
			fast++
			slow++
		}
	}
	return slow + 1
}

/*
两两交换链表中的节点
https://leetcode.cn/problems/swap-nodes-in-pairs/description/
输入：head = [1,2,3,4]
输出：[2,1,4,3]
*/
func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{0, head}
	tmp := newHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		n1 := tmp.Next
		n2 := tmp.Next.Next
		tmp.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		tmp = n1
	}
	return newHead.Next
}

/*
括号生成
https://leetcode.cn/problems/generate-parentheses/description/
示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
*/

func generateParenthesis(n int) []string {
	var gen func(left, right, n int, tmp string)
	var res []string
	gen = func(left, right, n int, tmp string) {
		if left == n && right == n {
			res = append(res, tmp)
			return
		}
		if left < n {
			gen(left+1, right, n, tmp+"(")
		}
		if right < left && right < n {
			gen(left, right+1, n, tmp+")")

		}
	}
	gen(0, 0, n, "")
	return res
}

/*
合并两个有序链表
https://leetcode.cn/problems/merge-two-sorted-lists/submissions/211224633/
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

/*
有效的括号
https://leetcode.cn/problems/valid-parentheses/description/
示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
*/
func isValid(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var queue []byte
	for _, v := range s {
		d := byte(v)
		if _, ok := m[d]; ok {
			queue = append(queue, d)
		} else {
			if len(queue) == 0 {
				return false
			}
			if m[queue[len(queue)-1]] == d {
				queue = queue[:len(queue)-1]
			} else {
				return false
			}
		}

	}
	return len(queue) == 0
}

/*
删除链表的倒数第 N 个结点
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	lenght := 0
	for node != nil {
		node = node.Next
		lenght++
	}
	removeIdx := lenght - n
	if removeIdx < 0 {
		return head
	}
	if removeIdx == 0 {
		head = head.Next
		return head
	}
	idx := 0
	node = head
	for idx < removeIdx-1 {
		node = node.Next
		idx++
	}
	node.Next = node.Next.Next
	return head

}

/*
 电话号码的字母组合
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

*/

func letterCombinations(digits string) []string {
	m := map[byte][]string{
		'1': {},
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var res []string
	if len(digits) == 0 {
		return res
	}
	var backTrack func(level int, tmp string)
	backTrack = func(level int, tmp string) {
		if level == len(digits) {
			res = append(res, tmp)
			return
		}
		for _, d := range m[digits[level]] {
			newTmp := tmp + d
			backTrack(level+1, newTmp)
		}

	}
	backTrack(0, "")
	return res
}

/*
三数之和
https://leetcode.cn/problems/3sum/description/
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	var res [][]int
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for {
			if l >= r {
				break
			}
			fmt.Println(l, r)
			sum := num + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{num, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
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
