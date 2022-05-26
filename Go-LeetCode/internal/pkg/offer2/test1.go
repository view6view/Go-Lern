package offer2

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
