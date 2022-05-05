package dp

import (
	"math"
	"sort"
	"strconv"
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

/**
跳跃游戏
https://leetcode-cn.com/problems/jump-game/
*/
func canJump(nums []int) bool {
	len := len(nums)
	if len == 1 {
		return true
	}
	if nums[0] <= 0 {
		return false
	}
	dp := make([]bool, len)
	dp[0] = true
	dp[1] = true
	var ok = 1
	var idx = 0
	for ok < len && idx <= ok {
		var canOk = idx + nums[idx]
		if canOk > ok {
			for i := ok + 1; i < len && i <= canOk; i++ {
				dp[i] = true
			}
			ok = canOk
		}
		idx++
	}
	return dp[len-1]
}

/**
最大子数组和
https://leetcode-cn.com/problems/maximum-subarray/
*/
func maxSubArray(nums []int) int {
	len := len(nums)
	pre := 0
	res := nums[0]
	for i := 0; i < len; i++ {
		pre = Max(pre+nums[i], nums[i])
		res = Max(res, pre)
	}
	return res
}

func maxSubArray1(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = Max(res, dp[i])
	}
	return res
}

/**
环形子数组的最大和
https://leetcode-cn.com/problems/maximum-sum-circular-subarray/
*/
func maxSubarraySumCircular(nums []int) int {
	total, maxSum, curMax, minSum, curMin := 0, nums[0], 0, nums[0], 0
	for _, num := range nums {
		curMax = Max(curMax+num, num)
		maxSum = Max(maxSum, curMax)
		curMin = Min(curMin+num, num)
		minSum = Min(minSum, curMin)
		total += num
	}
	if maxSum > 0 {
		return Max(maxSum, total-minSum)
	} else {
		return maxSum
	}
}

/**
乘积最大子数组
https://leetcode-cn.com/problems/maximum-product-subarray/
*/
func maxProduct(nums []int) int {
	max, iMax, iMin := math.MinInt32, 1, 1
	for _, num := range nums {
		// 如果当前num 小于0，则需要交换max 和 min
		if num < 0 {
			temp := iMax
			iMax = iMin
			iMin = temp
		}
		iMax = Max(iMax*num, num)
		iMin = Min(iMin*num, num)

		max = Max(max, iMax)
	}
	return max
}

/**
乘积为正数的最长子数组长度
https://leetcode-cn.com/problems/maximum-length-of-subarray-with-positive-product/
*/
func getMaxLen(nums []int) int {
	// 最大正数长度，前一个正数长度， 前一个负数长度
	maxLen, iMax, iMin := 0, 0, 0
	for _, num := range nums {
		if num == 0 {
			iMax = 0
			iMin = 0
			continue
		} else if num > 0 {
			if iMin != 0 {
				iMin++
			}
			iMax++
		} else {
			if iMax == 0 && iMin == 0 {
				iMin = 1
				continue
			} else {
				temp := iMax
				if iMin == 0 {
					iMax = 0
				} else {
					iMax = iMin + 1
				}
				iMin = temp + 1
			}
		}
		maxLen = Max(maxLen, iMax)
	}
	return maxLen
}

/**
最佳买卖股票时机含冷冻期
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
*/
func maxProfit4(prices []int) int {
	len := len(prices)

	// 定义dp
	dp := make([][]int, len)
	// 初始化dp
	dp[0] = make([]int, 3)
	// 不持股没卖出，最大利润
	dp[0][0] = 0
	// 持股，最大利润
	dp[0][1] = -prices[0]
	// 不持股卖出，最大利润
	dp[0][2] = 0

	for i := 1; i < len; i++ {
		dp[i] = make([]int, 3)
		// 状态转移
		dp[i][0] = Max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}

	return Max(dp[len-1][0], dp[len-1][2])
}

/**
买卖股票的最佳时机含手续费
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
*/
func maxProfit5(prices []int, fee int) int {
	len := len(prices)
	// 定义dp
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0]

	for i := 1; i < len; i++ {
		temp := dp[0]
		dp[0] = Max(dp[0], dp[1]+prices[i]-fee)
		dp[1] = Max(dp[1], temp-prices[i])
	}

	return dp[0]
}

/**
接雨水
https://leetcode-cn.com/problems/trapping-rain-water/
*/
func trap(height []int) int {
	len := len(height)

	maxLeft := make([]int, len)
	maxRight := make([]int, len)

	for i := 1; i < len; i++ {
		maxLeft[i] = Max(maxLeft[i-1], height[i-1])
	}

	for i := len - 2; i > -1; i-- {
		maxRight[i] = Max(maxRight[i+1], height[i+1])
	}

	res := 0

	for i := 1; i < len; i++ {
		minWall := Min(maxLeft[i], maxRight[i])
		if height[i] < minWall {
			res += minWall - height[i]
		}
	}
	return res
}

/**
解码方法
https://leetcode-cn.com/problems/decode-ways/
*/
func numDecodings(s string) int {
	len := len(s)
	if len == 0 || s[0] == '0' {
		return 0
	}

	// 定义dp
	dp := make([]int, len+1)
	// 初始化dp
	dp[0] = 1

	for i := 1; i <= len; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if i > 1 && s[i-2] != '0' {
			// 字符切片 左闭右开
			num, _ := strconv.ParseInt(s[i-2:i], 10, 10)
			if num < 27 && num > 0 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len]
}

/**
丑数 II
https://leetcode-cn.com/problems/ugly-number-ii/
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		num2, num3, num5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = Min(Min(num2, num3), num5)
		if dp[i] == num2 {
			p2++
		}
		if dp[i] == num3 {
			p3++
		}
		if dp[i] == num5 {
			p5++
		}
	}
	return dp[n]
}

/**
下降路径最小和
https://leetcode-cn.com/problems/minimum-falling-path-sum/
*/
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	dp := make([][]int, n)

	dp[0] = make([]int, n)
	copy(dp[0], matrix[0])

	for i := 1; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = Min(dp[i-1][0], dp[i-1][1]) + matrix[i][0]
		for j := 1; j < n-1; j++ {
			dp[i][j] = Min(dp[i-1][j], Min(dp[i-1][j-1], dp[i-1][j+1])) + matrix[i][j]
		}
		dp[i][n-1] = Min(dp[i-1][n-2], dp[i-1][n-1]) + matrix[i][n-1]
	}

	res := math.MaxInt

	for i := 0; i < n; i++ {
		res = Min(res, dp[n-1][i])
	}

	return res
}
