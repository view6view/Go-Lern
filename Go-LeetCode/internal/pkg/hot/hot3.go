package main

import (
	"container/list"
	"fmt"
	"strings"
)

/**
相交链表
https://leetcode.cn/problems/intersection-of-two-linked-lists/
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA := headA
	nodeB := headB
	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}
		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nodeA
}

/**
多数元素
https://leetcode.cn/problems/majority-element/
*/
func majorityElement(nums []int) int {
	candNum, count := nums[0], 0
	for _, num := range nums {
		if num == candNum {
			count++
		} else {
			count--
			if count == 0 {
				candNum = num
				count = 1
			}
		}
	}
	return candNum
}

/**
岛屿数量
https://leetcode.cn/problems/number-of-islands/
*/
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	var isValid func(i int, j int) bool
	isValid = func(i int, j int) bool {
		if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == '0' {
			return false
		}
		return true
	}

	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if isValid(i, j) {
			grid[i][j] = '0'
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
	}

	var bfs func(i int, j int)
	bfs = func(i int, j int) {
		list := list.New()
		list.PushBack([]int{i, j})
		for list.Len() > 0 {
			cur := list.Remove(list.Front()).([]int)
			if isValid(cur[0], cur[1]) {
				grid[cur[0]][cur[1]] = '0'
				list.PushBack([]int{cur[0] - 1, cur[1]})
				list.PushBack([]int{cur[0] + 1, cur[1]})
				list.PushBack([]int{cur[0], cur[1] - 1})
				list.PushBack([]int{cur[0], cur[1] + 1})
			}
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				bfs(i, j)
				count++
			}
		}
	}
	return count
}

/**
反转链表
https://leetcode.cn/problems/reverse-linked-list/
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := list.New()
	node := head
	for node != nil {
		list.PushBack(node)
		node = node.Next
	}

	head = list.Remove(list.Back()).(*ListNode)
	node = head
	for list.Len() > 0 {
		node.Next = list.Remove(list.Back()).(*ListNode)
		node = node.Next
	}
	node.Next = nil
	return head
}

/**
课程表
https://leetcode.cn/problems/course-schedule/
*/
// BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 对应课程所需要的前置条件数量
	needs := make([]int, numCourses)
	// 记录k（当前课程）可作为v（对应这些课程）的前置条件
	m := map[int][]int{}
	list := list.New()

	for _, prerequisite := range prerequisites {
		needs[prerequisite[0]]++
		m[prerequisite[1]] = append(m[prerequisite[1]], prerequisite[0])
	}

	for i, need := range needs {
		if need == 0 {
			list.PushBack(i)
		}
	}

	for list.Len() > 0 {
		cur := list.Remove(list.Front()).(int)
		numCourses--
		for _, i := range m[cur] {
			needs[i]--
			if needs[i] == 0 {
				list.PushBack(i)
			}
		}
	}
	return numCourses == 0
}

// DFS
func canFinish1(numCourses int, prerequisites [][]int) bool {
	// 记录对应i课程是否可以学习
	flags := make([]int, numCourses)
	// 记录k（当前课程）可作为v（对应这些课程）的前置条件
	m := map[int][]int{}

	var dfs func(i int) bool

	dfs = func(i int) bool {
		if flags[i] == -1 {
			return false
		}
		if flags[i] == 1 {
			return true
		}
		flags[i] = -1
		for _, course := range m[i] {
			if !dfs(course) {
				return false
			}
		}
		flags[i] = 1
		return true
	}

	for _, prerequisite := range prerequisites {
		m[prerequisite[1]] = append(m[prerequisite[1]], prerequisite[0])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

/**
数组中的第K个最大元素
https://leetcode.cn/problems/kth-largest-element-in-an-array/
*/
func findKthLargest(nums []int, k int) int {
	var swap = func(a int, b int) {
		temp := nums[a]
		nums[a] = nums[b]
		nums[b] = temp
	}

	var parition = func(low int, high int) int {
		pivot := nums[high]
		pointer := low
		for i := low; i < high; i++ {
			if nums[i] >= pivot {
				swap(i, pointer)
				pointer++
			}
		}
		swap(pointer, high)
		return pointer
	}

	target := k - 1
	left := 0
	right := len(nums) - 1

	for left <= right {
		idx := parition(left, right)
		if idx == target {
			return nums[idx]
		} else if idx > target {
			right = idx - 1
		} else {
			left = idx + 1
		}
	}
	return -1
}

/**
最大正方形
https://leetcode.cn/problems/maximal-square/
*/
func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])

	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}

	var min3 = func(a int, b int, c int) int {
		if b < a {
			a = b
		}
		if c < a {
			a = c
		}
		return a
	}

	var max = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxSide := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min3(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
				maxSide = max(maxSide, dp[i+1][j+1])
			}
		}
	}
	return maxSide * maxSide
}

/**
翻转二叉树
https://leetcode.cn/problems/invert-binary-tree/
*/
func invertTree(root *TreeNode) *TreeNode {
	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {
		if node != nil && (node.Left != nil || node.Right != nil) {
			left := node.Left
			node.Left = node.Right
			node.Right = left
			invert(node.Left)
			invert(node.Right)
		}
	}

	invert(root)
	return root
}

/**
回文链表
https://leetcode.cn/problems/palindrome-linked-list/
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢指针
	slow, fast := head, head
	pre := head
	var prePre *ListNode
	for fast != nil && fast.Next != nil {
		// 保存当前slow
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
		pre.Next = prePre
		prePre = pre
	}

	// 链表的节点个数为奇数
	if fast != nil {
		slow = slow.Next
	}
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		// 向前遍历
		pre = pre.Next
		// 向后遍历
		slow = slow.Next
	}
	return true
}

/**
二叉树的最近公共祖先
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果左边搜索结果为空，返回右子树搜索结果
	if left == nil {
		return right
	}
	// 如果右边搜索结果为空，返回左子树搜索结果
	if right == nil {
		return left
	}
	return root
}

/**
除自身以外数组的乘积
https://leetcode.cn/problems/product-of-array-except-self/
*/
func productExceptSelf(nums []int) []int {
	len := len(nums)
	res := make([]int, len)
	p := 1
	for i := 0; i < len; i++ {
		res[i] = p
		p *= nums[i]
	}
	p = 1
	for i := len - 1; i > 0; i-- {
		p *= nums[i]
		res[i-1] *= p
	}
	return res
}

/**
搜索二维矩阵 II
https://leetcode.cn/problems/search-a-2d-matrix-ii/submissions/
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

/**
移动零
https://leetcode.cn/problems/move-zeroes/
*/
func moveZeroes(nums []int) {
	idx := -1
	for i, num := range nums {
		if num != 0 {
			idx++
			temp := nums[idx]
			nums[idx] = nums[i]
			nums[i] = temp
		}
	}
}

/**
最长递增子序列
https://leetcode.cn/problems/longest-increasing-subsequence/
*/
func lengthOfLIS(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	res := 0
	for i, num := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < num {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		res = Max(res, dp[i])
	}
	return res + 1
}

/**
零钱兑换
https://leetcode.cn/problems/coin-change/
*/
func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := 1; i < max; i++ {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = Min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == max {
		return -1
	} else {
		return dp[amount]
	}
}

/**
打家劫舍 III
https://leetcode.cn/problems/house-robber-iii/
*/
func rob(root *TreeNode) int {
	m := map[*TreeNode]int{}

	fmt.Println(m[root])

	var inner func(root *TreeNode) int

	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	inner = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if m[node] != 0 {
			return m[node]
		}
		total := node.Val
		if node.Left != nil {
			total = total + inner(node.Left.Left) + inner(node.Left.Right)
		}
		if node.Right != nil {
			total = total + inner(node.Right.Left) + inner(node.Right.Right)
		}
		total = max(total, inner(node.Left)+inner(node.Right))
		m[node] = total
		return total
	}
	return inner(root)
}

/**
比特位计数
https://leetcode.cn/problems/counting-bits/
*/
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&1 == 1 { // 奇数
			res[i] = res[i-1] + 1
		} else { // 偶数
			res[i] = res[i/2]
		}
	}
	return res
}

/**
前 K 个高频元素
https://leetcode.cn/problems/top-k-frequent-elements/
*/
func topKFrequent(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}

	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	// 使用 桶排序 进行频次排序
	buckets := make([][]int, l)
	for num, cnt := range m {
		if buckets[cnt] == nil {
			buckets[cnt] = make([]int, 0)
		}
		buckets[cnt] = append(buckets[cnt], num)
	}

	res := make([]int, k)
	idx := 0
	for i := l - 1; i >= 0; i-- {
		// 空桶，跳过
		if buckets[i] == nil {
			continue
		}
		for _, num := range buckets[i] {
			res[idx] = num
			idx++
		}
		if idx == k {
			break
		}
	}
	return res
}

/**
字符串解码
https://leetcode.cn/problems/decode-string/
*/
func decodeString(s string) string {
	l := len(s)

	var isNum = func(c byte) bool {
		if c <= '9' && c >= '0' {
			return true
		}
		return false
	}
	var getStrAndIdx func(i int) (string, int)
	getStrAndIdx = func(i int) (string, int) {
		cnt := s[i] - '0'
		begin := i + 1
		for isNum(s[begin]) {
			cnt = cnt*10 + (s[begin] - '0')
			begin++
		}
		// [ 符号的下一位字符
		begin++
		builder := strings.Builder{}
		for ; s[begin] != ']'; begin++ {
			if isNum(s[begin]) {
				str, idx := getStrAndIdx(begin)
				builder.WriteString(str)
				begin = idx
			} else {
				builder.WriteByte(s[begin])
			}

		}
		cur := builder.String()
		// 当前 builder 已经存在了1种，所以只需要判断大于1
		for cnt > 1 {
			builder.WriteString(cur)
			cnt--
		}
		return builder.String(), begin
	}

	builder := strings.Builder{}
	for i := 0; i < l; i++ {
		c := s[i]
		if c <= '9' && c >= '0' {
			str, idx := getStrAndIdx(i)
			builder.WriteString(str)
			i = idx
		} else {
			builder.WriteByte(c)
		}
	}
	return builder.String()
}
