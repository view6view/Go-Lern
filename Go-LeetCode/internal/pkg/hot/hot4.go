package main

import (
	"container/list"
	"sort"
)

/**
根据身高重建队列
https://leetcode.cn/problems/queue-reconstruction-by-height/
*/
func reconstructQueue(people [][]int) [][]int {
	// 按照身高降序，如果身高相同，按照前面人个数升序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	res := make([][]int, len(people))
	for idx, person := range people {
		// idx代表当前res中人的个数，person[1]代表person前面身高大于person的个数
		// 如果 idx > person[1] ，就执行下面代码将person插队到person[1]位置
		if idx > person[1] {
			for i := idx; i > person[1]; i-- {
				res[i] = res[i-1]
			}
			res[person[1]] = person
		} else {
			// 如果 idx <= person[1] ，就将person插队到末尾
			res[idx] = person
		}
	}
	return res
}

/**
找到字符串中所有字母异位词
https://leetcode.cn/problems/find-all-anagrams-in-a-string/
*/
func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	m := map[byte]int{}
	for i := 0; i < pLen; i++ {
		m[p[i]]++
	}

	window := map[byte]int{}
	left, right := 0, 0
	// 需要符合的有效字符个数
	valid := pLen

	res := make([]int, 0)

	for right < sLen {
		rChar := s[right]
		if m[rChar] != 0 {
			window[rChar]++
			// 只有window中字符数量不超过m，才算一次有效包含
			if window[rChar] <= m[rChar] {
				valid--
			}
		}
		for valid == 0 {
			if right-left+1 == pLen {
				res = append(res, left)
			}
			lChar := s[left]
			if m[lChar] != 0 {
				window[lChar]--
				// 有效字符个数减少
				if window[lChar] < m[lChar] {
					valid++
				}
			}
			left++
		}
		right++
	}
	return res
}

/**
找到所有数组中消失的数字
https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/
*/
func findDisappearedNumbers(nums []int) []int {
	len := len(nums)
	flags := make([]bool, len+1)
	for _, num := range nums {
		flags[num] = true
	}
	res := make([]int, 0)
	for i := 1; i <= len; i++ {
		if !flags[i] {
			res = append(res, i)
		}
	}
	return res
}

/**
汉明距离
https://leetcode.cn/problems/hamming-distance/
*/
func hammingDistance(x int, y int) int {
	cnt := 0
	p := x ^ y
	for p > 0 {
		if p&1 == 1 {
			cnt++
		}
		p >>= 1
	}
	return cnt
}

/**
把二叉搜索树转换为累加树
https://leetcode.cn/problems/convert-bst-to-greater-tree/
*/
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 用于递归过程中记录前一个节点的值
	pre := 0

	// 递归遍历，遍历顺序和中序相反
	var inner func(node *TreeNode)
	inner = func(node *TreeNode) {
		if node.Right != nil {
			inner(node.Right)
		}
		node.Val += pre
		pre = node.Val
		if node.Left != nil {
			inner(node.Left)
		}
	}

	inner(root)

	return root
}

/**
非递归，记录中序遍历的顺序
*/
func convertBST1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 用于存储中序遍历顺序
	l1 := list.New()
	// 用于非递归实现中序遍历
	l2 := list.New()

	// 中序遍历
	temp := root
	for temp != nil || l2.Len() > 0 {
		if temp != nil {
			l2.PushBack(temp)
			temp = temp.Left
		} else {
			temp = l2.Remove(l2.Back()).(*TreeNode)
			l1.PushBack(temp)
			temp = temp.Right
		}
	}

	pre := l1.Remove(l1.Back()).(*TreeNode).Val

	for l1.Len() > 0 {
		temp = l1.Remove(l1.Back()).(*TreeNode)
		temp.Val += pre
		pre = temp.Val
	}

	return root
}

/**
二叉树的直径
https://leetcode.cn/problems/diameter-of-binary-tree/
*/
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 用于记录可能存在不经过根节点的情况
	maxNotForRoot := 0

	var getMaxHeight func(node *TreeNode) int
	getMaxHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := getMaxHeight(node.Left)
		r := getMaxHeight(node.Right)
		maxNotForRoot = max(maxNotForRoot, l+r)
		maxHeight := max(l, r) + 1
		return maxHeight
	}

	maxForRoot := getMaxHeight(root.Left) + getMaxHeight(root.Right)
	return max(maxForRoot, maxNotForRoot)
}

/**
和为 K 的子数组
https://leetcode.cn/problems/subarray-sum-equals-k/
*/
func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	preSum, cnt := 0, 0
	for _, num := range nums {
		preSum += num
		// 先获得前缀和为 preSum - k 的个数，加到计数变量里
		// 获得和为preSum - k的数组前缀，可以得到满足当前前缀的子数组个数，因为有负数，可能存在多个
		cnt += m[preSum-k]
		// 然后维护 preSumFreq 的定义
		m[preSum]++
	}
	return cnt
}

/**
最短无序连续子数组
https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/
*/
func findUnsortedSubarray(nums []int) int {
	len := len(nums)
	max, min := nums[0], nums[len-1]
	begin, end := 0, -1
	for left, right := 0, len-1; left < len; left, right = left+1, right-1 {
		// 从左往右去更新最大值和下标
		if nums[left] < max {
			end = left
		} else {
			max = nums[left]
		}

		// 从右往左去更新最小值和下标
		if nums[right] > min {
			begin = right
		} else {
			min = nums[right]
		}
	}
	return end - begin + 1
}

/**
合并二叉树
https://leetcode.cn/problems/merge-two-binary-trees/
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 深度遍历去构造树
	var dfs func(root1 *TreeNode, root2 *TreeNode) *TreeNode
	dfs = func(root1 *TreeNode, root2 *TreeNode) *TreeNode {
		if root1 == nil {
			return root2
		}
		if root2 == nil {
			return root1
		}
		root1.Val += root2.Val
		root1.Left = dfs(root1.Left, root2.Left)
		root1.Right = dfs(root1.Right, root2.Right)
		return root1
	}
	return dfs(root1, root2)
}

/**
任务调度器
https://leetcode.cn/problems/task-scheduler/
*/
func leastInterval(tasks []byte, n int) int {
	buckets := make([]int, 26)
	for _, task := range tasks {
		buckets[task-'A']++
	}
	sort.Ints(buckets)

	maxCount := 1
	for i := 25; i > 0; i-- {
		if buckets[i] == buckets[i-1] {
			maxCount++
		} else {
			break
		}
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// n为0的时候，len(tasks)是最佳情况
	return max(len(tasks), (buckets[25]-1)*(n+1)+maxCount)
}

/**
回文子串
https://leetcode.cn/problems/palindromic-substrings/
*/
func countSubstrings(s string) int {
	len := len(s)
	dp := make([][]bool, len)
	cnt := 0
	for i := 0; i < len; i++ {
		dp[i] = make([]bool, len)
		for j := 0; j <= i; j++ {
			if s[j] == s[i] && (i-j < 2 || dp[j+1][i-1]) {
				cnt++
				dp[j][i] = true
			}
		}
	}
	return cnt
}

/**
每日温度
https://leetcode.cn/problems/daily-temperatures/
*/
func dailyTemperatures(temperatures []int) []int {
	len := len(temperatures)
	res := make([]int, len)
	// 递减栈实现
	l := list.New()
	for i, temperature := range temperatures {
		for l.Len() > 0 && temperature > temperatures[l.Back().Value.(int)] {
			preIdx := l.Remove(l.Back()).(int)
			res[preIdx] = i - preIdx
		}
		l.PushBack(i)
	}
	return res
}
