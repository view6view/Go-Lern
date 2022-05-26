package main

import (
	"Go-LeetCode/internal/pkg/data"
	"container/list"
	"math"
)

/**
删除有序数组中的重复项
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
*/
func removeDuplicates(nums []int) int {
	cnt := 0
	len := len(nums)
	for i := 0; i < len; {
		nums[cnt] = nums[i]
		cnt++
		j := i + 1
		for ; j < len && nums[j] == nums[i]; j++ {
		}
		i = j
	}
	return cnt
}

/**
单词搜索
https://leetcode.cn/problems/word-search/
*/
func exist(board [][]byte, word string) bool {
	var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	rows := len(board)
	cols := len(board[0])
	len := len(word)
	visited := make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}

	// 声明函数，判断当前点是否在区域内
	var isArea func(x int, y int) bool
	isArea = func(x int, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}
	// 声明函数，深度优先遍历
	var dfs func(x int, y int, begin int) bool
	dfs = func(x int, y int, begin int) bool {
		if begin == len-1 {
			return board[x][y] == word[begin]
		}
		if board[x][y] == word[begin] {
			visited[x][y] = true
			for _, direction := range directions {
				newX := x + direction[0]
				newY := y + direction[1]
				if isArea(newX, newY) && !visited[newX][newY] {
					if dfs(newX, newY, begin+1) {
						return true
					}
				}
			}
			visited[x][y] = false
		}
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

/**
二叉树展开为链表
https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	queue := list.New()
	stack := list.New()

	node := root

	// 利用栈前序遍历
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			queue.PushBack(node)
			node = node.Left
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			node = node.Right
		}
	}

	// 对队列进行封装
	root = queue.Remove(queue.Front()).(*TreeNode)
	node = root
	for queue.Len() > 0 {
		node.Left = nil
		node.Right = queue.Remove(queue.Front()).(*TreeNode)
		node = node.Right
	}
}

/**
最长连续序列
https://leetcode.cn/problems/longest-consecutive-sequence/
*/
func longestConsecutive(nums []int) int {
	type void struct{}
	var member void
	set := make(map[int]void, len(nums))
	res := 0
	for _, num := range nums {
		small := num - 1
		big := num + 1
		for _, exists := set[small]; exists; _, exists = set[small] {
			small--
		}
		for _, exists := set[big]; exists; _, exists = set[big] {
			big++
		}
		res = Max(res, big-small-1)
		_, exists := set[num]
		if !exists {
			set[num] = member
		}
	}
	return res
}

/**
使用位图，没通过
*/
func longestConsecutive1(nums []int) int {
	max, min := math.MinInt, math.MaxInt
	for _, num := range nums {
		max = Max(max, num)
		min = Min(min, num)
	}
	bitMap := data.CreateNewBitMap(max - min)
	for _, num := range nums {
		bitMap.Add(num)
	}

	var res int
	var curLen int
	for i := 0; i < bitMap.Max+1; i++ {
		if bitMap.Exist(i) {
			curLen++
			res = Max(res, curLen)
		} else {
			curLen = 0
		}
	}
	return res
}

/**
出现一次的数字
https://leetcode.cn/problems/single-number/
*/
func singleNumber(nums []int) (res int) {
	for _, num := range nums {
		res ^= num
	}
	return res
}

/**
环形链表 II
https://leetcode.cn/problems/linked-list-cycle-ii/
*/
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for true {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
