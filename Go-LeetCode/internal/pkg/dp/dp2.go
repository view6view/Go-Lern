package dp

import (
	"sort"
)

/**
买卖股票的最佳时机 II
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
*/
func maxProfit2(prices []int) int {
	len := len(prices)

	dp := make([][]int, len)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len; i++ {
		if dp[i-1][0] == 0 && prices[i] < prices[i-1] {
			// 找到第一个低谷为止
			dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = Max(dp[i-1][1], -prices[i])
		} else if prices[i] < prices[i-1] {
			// 找到下一个低谷为止
			dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = Max(dp[i-1][1], -prices[i])
		} else {
			// 计算当前相对于上一个谷底的盈利
			dp[i][0] = Max(dp[i-1][0], dp[i-1][0]+dp[i-1][1]+prices[i])
			dp[i][1] = Min(dp[i-1][1], -prices[i])
		}
	}
	return dp[len-1][0]
}

/**
贪心法
今天相对于明天有新增就加上
*/
func maxProfit3(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		temp := prices[i] - prices[i-1]
		if temp > 0 {
			res += temp
		}
	}
	return res
}

/**
等差数列划分
https://leetcode-cn.com/problems/arithmetic-slices/
*/
func numberOfArithmeticSlices(nums []int) int {
	len := len(nums)
	if len < 3 {
		return 0
	}

	// 定义dp
	dp := make([]int, len)
	// 初始化dp
	dp[0] = 0
	dp[1] = 0

	res := 0

	for i := 2; i < len; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			res += dp[i]
		}
	}

	return res
}

/**
杨辉三角
https://leetcode-cn.com/problems/pascals-triangle/
*/
func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		arr := make([]int, i)
		arr[0] = 1
		if i > 1 {
			if i > 2 {
				for j := 1; j < i-1; j++ {
					arr[j] = dp[i-2][j-1] + dp[i-2][j]
				}
			}
			arr[i-1] = 1
		}
		dp[i-1] = arr
	}
	return dp
}

/**
杨辉三角 II
https://leetcode-cn.com/problems/pascals-triangle-ii/
*/
func getRow(rowIndex int) []int {
	// 定义dp
	dp := make([]int, rowIndex+1)
	// 初始化dp
	dp[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp
}

/**
不同路径
https://leetcode-cn.com/problems/unique-paths/
*/
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}

/**
不同路径 II
https://leetcode-cn.com/problems/unique-paths-ii/
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// 定义 dp
	dp := make([][]int, m)

	// 初始化dp
	dp[0] = make([]int, n)
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] != 1 {
			dp[0][i] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] != 1 && dp[i-1][0] == 1 {
			dp[i][0] = 1
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]
}

/**
整数拆分
https://leetcode-cn.com/problems/integer-break/
*/
func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}
	res := 1
	for n >= 5 {
		res *= 3
		n -= 3
	}
	res *= n
	return res
}

/**
不同的二叉搜索树
https://leetcode-cn.com/problems/unique-binary-search-trees/
*/
func numTrees(n int) int {
	if n < 3 {
		return n
	}
	// 定义dp
	dp := make([]int, n+1)
	// 初始化dp
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

/**
打家劫舍
https://leetcode-cn.com/problems/house-robber/
*/
func rob(nums []int) int {
	len := len(nums)
	if len == 1 {
		return nums[0]
	}
	if len == 2 {
		return Max(nums[0], nums[1])
	}
	// 定义dp
	dp := make([]int, len)
	// 初始化dp
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < len; i++ {
		dp[i] = Max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len-1]
}

/**
打家劫舍2
https://leetcode-cn.com/problems/house-robber-ii/
*/
func rob2(nums []int) int {
	len := len(nums)
	if len == 1 {
		return nums[0]
	}
	if len == 2 {
		return Max(nums[0], nums[1])
	}
	return Max(myRob2(nums, 0, len-2), myRob2(nums, 1, len-1))
}

func myRob2(nums []int, begin int, end int) int {
	pre := nums[begin]
	next := Max(nums[begin], nums[begin+1])
	for i := begin + 2; i <= end; i++ {
		temp := next
		next = Max(nums[i]+pre, next)
		pre = temp
	}
	return next
}

/**
删除并获得点数
https://leetcode-cn.com/problems/delete-and-earn/
*/
func deleteAndEarn(nums []int) int {
	len := len(nums)
	if len == 1 {
		return nums[0]
	}
	// 排序
	sort.Ints(nums)
	max := nums[len-1]

	// 定义dp
	dp := make([]int, max+1)

	// 初始化dp
	start := 0
	end := 1
	temp := nums[start]
	for end < len && nums[end] == nums[start] {
		temp += nums[start]
		end++
	}
	dp[nums[start]] = temp

	// 遍历dp
	idx := nums[start] + 1
	start = end
	end++
	for ; idx <= max; idx++ {
		if nums[start] == idx {
			temp = nums[start]
			for end < len && nums[end] == nums[start] {
				temp += nums[start]
				end++
			}
			dp[nums[start]] = Max(dp[nums[start]-1], dp[nums[start]-2]+temp)
			if end == len {
				return dp[max]
			}
			start = end
			end++
		} else {
			dp[idx] = Max(dp[idx-1], dp[idx-2])
		}
	}
	return dp[max]
}
