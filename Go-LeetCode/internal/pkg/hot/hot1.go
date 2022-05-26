package main

import "sort"

/**
两数之和
https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	idxMap := map[int]int{}
	for i, num := range nums {
		if v, ok := idxMap[target-num]; ok {
			return []int{v, i}
		}
		idxMap[num] = i
	}
	return nil
}

/**
两数相加
https://leetcode-cn.com/problems/add-two-numbers/
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	//定义一个尾结点，或者可以理解为临时节点
	var tail *ListNode
	//余数carry
	carry := 0
	//依次遍历两个链表，只要元素不为空就进行下一步
	for l1 != nil || l2 != nil {
		//定义两个变量存储各个节点的值
		n1, n2 := 0, 0
		//从第一个链表开始
		if l1 != nil {
			//把每个节点的值赋给n1
			n1 = l1.Val
			//节点后移
			l1 = l1.Next
		}
		//l2同上
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//此时是两个链表第一个元素的和 + 余数
		sum := n1 + n2 + carry
		//sum%10是节点的当前值，如果是10,取余后当前节点值为0，sum/10是求十位的那个数
		sum, carry = sum%10, sum/10
		//此时申请一个新的链表存储两个链表的和
		if head == nil {
			//申请新的链表
			head = &ListNode{Val: sum}
			//这一步是为了保持头结点不变的情况下指针可以右移，所以说tail相当于临时节点，理解成尾节点也可以，因
			//为此时新链表中只有一个节点，所以头结点和尾结点都指向同一个元素。
			tail = head
		} else {
			//第二个节点后开始逐渐往尾结点增加元素
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	//把最后一位的余数加到链表最后。
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

/**
下一个排列
https://leetcode-cn.com/problems/next-permutation/
*/
func nextPermutation(nums []int) {
	len := len(nums)

	if len <= 1 {
		return
	}

	i, j, k := len-2, len-1, len-1

	// 找到num[i] < nums[j]
	for j > 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if j > 0 {
		// 从最后一个数向前开始找到第一个大于 nums[i]的数据
		for nums[i] >= nums[k] {
			k--
		}
		// 交换nums[i] nums[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 反转nums[j:end]
	for begin, end := j, len-1; begin < end; begin, end = begin+1, end-1 {
		// 交换
		nums[begin], nums[end] = nums[end], nums[begin]
	}
}

/**
https://leetcode.cn/problems/search-in-rotated-sorted-array/
搜索旋转排序数组
*/
func search(nums []int, target int) int {
	len := len(nums)
	if len == 0 {
		return -1
	}
	left, right := 0, len-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] < nums[right] { // 右边有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // 左边有序
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

/**
在排序数组中查找元素的第一个和最后一个位置
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
*/
func searchRange(nums []int, target int) []int {
	len := len(nums)
	left, right := 0, len-1
	if len == 0 || target < nums[left] || target > nums[right] {
		return []int{-1, -1}
	}
	for left <= right {
		mid := (left + right) / 2

		// 找到值
		if target == nums[mid] {
			left = mid - 1
			right = mid + 1
			for left > -1 && nums[left] == nums[mid] {
				left--
			}
			for right < len && nums[right] == nums[mid] {
				right++
			}
			return []int{left + 1, right - 1}
		}

		if target < nums[mid] { // 左边
			right = mid - 1
			for right >= left && nums[right] == nums[mid] {
				right--
			}
		} else { // 右边
			left = mid + 1
			for left <= right && nums[left] == nums[mid] {
				left++
			}
		}
	}
	return []int{-1, -1}
}

/**
两两交换链表中的节点
https://leetcode.cn/problems/swap-nodes-in-pairs/
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{0, head}
	iterator := newHead

	for iterator.Next != nil && iterator.Next.Next != nil {
		temp := iterator.Next
		iterator.Next = iterator.Next.Next
		temp.Next = iterator.Next.Next
		iterator.Next.Next = temp
		iterator = temp
	}
	return newHead.Next
}

/**
组合总和
https://leetcode.cn/problems/combination-sum/
*/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	var path []int

	var dfs func(candidates []int, begin int, target int)
	dfs = func(candidates []int, begin int, target int) {
		if target < 0 {
			return
		}

		if target == 0 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			res = append(res, newPath)
			return
		}

		for i := begin; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(candidates, i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}

	dfs(candidates, 0, target)
	return res
}

/**
全排列！
https://leetcode.cn/problems/permutations/
*/
func permute(nums []int) [][]int {
	len := len(nums)
	var res [][]int
	// 临时数组
	var path []int
	// 记录当前遍历了的数组
	pathFlag := make([]bool, len)

	var dfs func(size int)
	dfs = func(size int) {
		if size == len {
			newPath := make([]int, len)
			copy(newPath, path)
			res = append(res, newPath)
		} else {
			for i := 0; i < len; i++ {
				if pathFlag[i] {
					continue
				} else {
					path = append(path, nums[i])
					pathFlag[i] = true
					dfs(size + 1)
					path = path[:size]
					pathFlag[i] = false
				}
			}
		}
	}
	dfs(0)
	return res
}

/**
旋转图像!
https://leetcode.cn/problems/rotate-image/
*/
func rotate(matrix [][]int) {
	len := len(matrix)

	// 上下部分进行交换
	for i := 0; i < len/2; i++ {
		for j := 0; j < len; j++ {
			i2 := len - i + 1
			temp := matrix[i][j]
			matrix[i][j] = matrix[i2][j]
			matrix[i2][j] = temp
		}
	}
	// 以 \ 为对角线，右上三角与左下三角进行交换
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}

/**
字母异位词分组
https://leetcode.cn/problems/group-anagrams/submissions/
*/
func groupAnagrams(strs []string) (ans [][]string) {
	m := map[string][]string{}

	for _, str := range strs {
		sa := []byte(str)
		// 排序
		sort.Slice(sa, func(i, j int) bool {
			return sa[i] < sa[j]
		})

		sign := string(sa)
		m[sign] = append(m[sign], str)
	}

	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}

/**
换一种更高效的sign算法
*/
func groupAnagrams1(strs []string) (ans [][]string) {
	// 签名算法
	sign := func(str string) string {
		var b [26]byte
		for _, s := range str {
			b[s-'a']++
		}
		return string(b[:])
	}

	m := map[string][]string{}
	for _, str := range strs {
		s := sign(str)
		m[s] = append(m[s], str)
	}
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}

/**
合并区间
https://leetcode.cn/problems/merge-intervals/
*/
func merge(intervals [][]int) (res [][]int) {
	// 进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		peek := intervals[i]
		if temp[1] >= peek[0] {
			temp[1] = Max(temp[1], peek[1])
		} else {
			res = append(res, temp)
			temp = peek
		}
	}
	res = append(res, temp)
	return res
}

/**
最小路径和
https://leetcode.cn/problems/minimum-path-sum/
*/
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += Min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

/**
颜色分类
https://leetcode.cn/problems/sort-colors/
*/
func sortColors(nums []int) {
	idx0, idx1 := 0, 0
	for i, num := range nums {
		nums[i] = 2
		if num < 2 {
			nums[idx1] = 1
			idx1++
			if num < 1 {
				nums[idx0] = 0
				idx0++
			}
		}
	}
}

/**
从前序与中序遍历序列构造二叉树！
https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/
func buildTree(preorder []int, inorder []int) (root *TreeNode) {
	m := map[int]int{}
	for i, num := range inorder {
		m[num] = i
	}

	var helper func(preorder []int, pStart int, pEnd int, inorder []int, iStart int, m map[int]int) (root *TreeNode)

	helper = func(preorder []int, pStart int, pEnd int, inorder []int, iStart int, m map[int]int) (root *TreeNode) {
		if pStart == pEnd {
			return nil
		}
		rootVal := preorder[pStart]
		root = &TreeNode{rootVal, nil, nil}
		// 当前root在中序遍历中的下标
		iIdx := m[rootVal]
		// 当前节点的左子树结束节点位置：pStart + 当前root的左子节点个数 + 1
		newPEnd := pStart + (iIdx - iStart) + 1
		root.Left = helper(preorder, pStart+1, newPEnd, inorder, iStart, m)
		root.Right = helper(preorder, newPEnd, pEnd, inorder, iIdx+1, m)
		return root
	}
	return helper(preorder, 0, len(preorder), inorder, 0, m)
}
