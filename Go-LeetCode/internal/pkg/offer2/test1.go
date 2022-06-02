package offer2

import (
	"container/list"
	"math"
	"strings"
)

/**
矩阵中的路径
https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/
*/
func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	k := len(word) - 1

	var dfs func(row, col int, cnt int) bool

	dfs = func(row, col int, idx int) bool {
		// 如果当前遍历坐标超出范围 或者 word[idx] 不等于 board[row][col]，则可以返回false
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[idx] {
			return false
		}
		// 匹配到最后一个字符符合，可以返回true
		if idx == k {
			return true
		}
		// 先将当前这个点设置为一个 不会匹配的字符，防止递归遍历的时候判断异常
		board[row][col] = '\n'
		nextIdx := idx + 1
		var res bool = dfs(row-1, col, nextIdx) ||
			dfs(row+1, col, nextIdx) ||
			dfs(row, col-1, nextIdx) ||
			dfs(row, col+1, nextIdx)
		// 将当前点还原，通过第一个 if 已经判断 board[row][col] == word[idx]，所以递归结束就还原
		board[row][col] = word[idx]
		return res
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
机器人的运动范围
https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
*/
func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	var bitSum = func(n int) int {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		return sum
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m || j >= n || visited[i][j] || bitSum(i)+bitSum(j) > k {
			return 0
		}
		visited[i][j] = true
		return 1 + dfs(i+1, j) + dfs(i, j+1)
	}
	return dfs(0, 0)
}

/**
剪绳子 II
https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/
*/
func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	MOD := 1000000007

	var getMax func(num int) int
	getMax = func(num int) int {
		if num <= 4 {
			return num
		} else {
			return 3 * getMax(num-3) % MOD
		}
	}
	return getMax(n)
}

/**
数值的整数次方
https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := 1.0
	// 如果负数，取反
	if n < 0 {
		x = 1 / x
		n = -n
	}
	// 快速幂思想
	for n > 0 {
		// 如果当前为奇数次，则乘上x
		if n&1 == 1 {
			res *= x
		}
		x *= x
		// n 左移一位
		n >>= 1
	}
	return res
}

/**
表示数值的字符串
https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
*/
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	// 标记是否 数值， 小数点， e或者E
	isNum, isDot, iseOrE := false, false, false
	for i := range s {
		c := s[i]
		if c >= '0' && c <= '9' {
			isNum = true
		} else if c == '.' {
			// 小数点不能重复出现，也不能同时出现小数点 和 e、E
			if isDot || iseOrE {
				return false
			}
			isDot = true
		} else if c == 'e' || c == 'E' {
			// e 和 E前面必须有数，并且不能重复出现e
			if !isNum || iseOrE {
				return false
			}
			iseOrE = true
			// 重置isNum，因为e或E之后也必须接上整数，防止出现 123e或者123e+的非法情况
			isNum = false
		} else if c == '-' || c == '+' {
			// 正负号只可能出现在第一个位置，或者出现在‘e’或'E'的后面一个位置
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
	}
	return isNum
}

/**
树的子结构
https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	// node1 必须 包含 node2
	var isMath func(node1, node2 *TreeNode) bool
	isMath = func(node1, node2 *TreeNode) bool {
		if node1 == nil || node2 == nil {
			if node1 == nil && node2 == nil {
				return true
			}
			if node1 == nil {
				return false
			} else {
				return true
			}
		}
		return node1.Val == node2.Val &&
			isMath(node1.Left, node2.Left) &&
			isMath(node1.Right, node2.Right)
	}

	var preOrder func(node *TreeNode) bool
	preOrder = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if node.Val == B.Val {
			if isMath(node.Left, B.Left) &&
				isMath(node.Right, B.Right) {
				return true
			}
		}
		return preOrder(node.Left) || preOrder(node.Right)
	}
	return preOrder(A)
}

/**
栈的压入、弹出序列
https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
*/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()
	len := len(pushed)
	idx1 := 0
	idx2 := 0
	for idx1 < len {
		if pushed[idx1] == popped[idx2] {
			idx2++
			for stack.Len() > 0 && stack.Back().Value.(int) == popped[idx2] {
				stack.Remove(stack.Back())
				idx2++
			}
		} else {
			stack.PushBack(pushed[idx1])
		}
		stack.PushBack(pushed[idx1])
		idx1++
	}
	for idx2 < len {
		if stack.Remove(stack.Back()).(int) == popped[idx2] {
			idx2++
		} else {
			return false
		}
	}
	return true
}

/**
用数组实现
*/
func validateStackSequences1(pushed []int, popped []int) bool {
	len := len(pushed)
	stack := make([]int, len)
	// 栈顶指针
	idx := -1
	// 入栈数组指针
	idx1 := 0
	// 出栈数组指针
	idx2 := 0
	for idx1 < len {
		if pushed[idx1] == popped[idx2] {
			idx2++
			for idx > -1 && stack[idx] == popped[idx2] {
				// 弹出元素
				idx--
				idx2++
			}
		} else {
			idx++
			stack[idx] = pushed[idx1]
		}
		idx1++
	}
	for idx2 < len {
		// 弹出元素对比
		if stack[idx] == popped[idx2] {
			idx--
			idx2++
		} else {
			return false
		}
	}
	return true
}

/**
从上到下打印二叉树
https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
*/
func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		front := l.Remove(l.Front()).(*TreeNode)
		res = append(res, front.Val)
		if front.Left != nil {
			l.PushBack(front.Left)
		}
		if front.Right != nil {
			l.PushBack(front.Right)
		}
	}
	return res
}

/**
开始 后进
第一层 前出(lr),后进(先l后r)
第二层 后出(rl),前进(先r后l)
第三层 前出(lr),后出(先l后r)
https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
*/
func levelOrder3(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	deque := list.New()
	// 定义是否式left to right
	lr := true

	packageLevel := func(curSize int) int {
		level := make([]int, curSize)
		if lr {
			for i := 0; i < curSize; i++ {
				treeNode := deque.Remove(deque.Front()).(*TreeNode)
				level[i] = treeNode.Val
				if treeNode.Left != nil {
					deque.PushBack(treeNode.Left)
				}
				if treeNode.Right != nil {
					deque.PushBack(treeNode.Right)
				}
			}
		} else {
			for i := 0; i < curSize; i++ {
				treeNode := deque.Remove(deque.Back()).(*TreeNode)
				level[i] = treeNode.Val
				if treeNode.Right != nil {
					deque.PushFront(treeNode.Right)
				}
				if treeNode.Left != nil {
					deque.PushFront(treeNode.Left)
				}
			}
		}
		res = append(res, level)
		return deque.Len()
	}

	deque.PushBack(root)
	size := 1
	for size > 0 {
		size = packageLevel(size)
		lr = !lr
	}
	return res
}

/**
二叉搜索树的后序遍历序列!
https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
*/
func verifyPostorder(postorder []int) bool {
	end := len(postorder) - 1
	stack := list.New()
	pre := math.MaxInt
	for i := end; i > -1; i-- {
		if postorder[i] > pre {
			return false
		}
		for stack.Len() > 0 && stack.Back().Value.(int) > postorder[i] {
			pre = stack.Remove(stack.Back()).(int)
		}
		stack.PushBack(postorder[i])
	}
	return true
}

/**
二叉树中和为某一值的路径
https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
*/
func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var dfs func(node *TreeNode, preTotal int, preArr []int)

	dfs = func(node *TreeNode, preTotal int, preArr []int) {
		preArr = append(preArr, node.Val)
		preTotal += node.Val
		if node.Left == nil && node.Right == nil {
			if preTotal == target {
				res = append(res, append([]int(nil), preArr...))
			}
		} else {
			if node.Left != nil {
				dfs(node.Left, preTotal, preArr)
			}
			if node.Right != nil {
				dfs(node.Right, preTotal, preArr)
			}
		}
	}
	dfs(root, 0, []int{})
	return res
}
