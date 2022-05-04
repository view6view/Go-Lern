package dp

import "math"

/*
最长回文子串
https://leetcode-cn.com/problems/longest-palindromic-substring/
*/
func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}

	// 定义dp数组
	dp := make([][]bool, len)
	for i := range dp {
		dp[i] = make([]bool, len)
	}
	// 初始化dp数组
	dp[0][0] = true
	begin := 0
	maxLen := 1

	for right := 1; right < len; right++ {
		dp[right][right] = true
		for left := 0; left < right; left++ {
			if s[left] != s[right] {
				dp[left][right] = false
			} else {
				// right - left < 3 且两端字符像相等，那么就是回文
				if right-left < 3 {
					dp[left][right] = true
				} else {
					// 划分子问题
					dp[left][right] = dp[left+1][right-1]
				}
				// 判断是否为回文字符串，更新
				if dp[left][right] {
					curLen := right - left + 1
					if curLen > maxLen {
						maxLen = curLen
						begin = left
					}
				}
			}
		}
	}
	// 截取字符串返回结果
	return s[begin : begin+maxLen]
}

/**
斐波那契数
https://leetcode-cn.com/problems/fibonacci-number/
*/
func fib(n int) int {
	if n < 2 {
		return n
	}

	// 定义dp
	dp := make([]int, n+1)
	// 初始化dp
	dp[0] = 0
	dp[1] = 1

	// 遍历背包
	for i := 2; i <= n; i++ {
		// 状态转换
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/**
第 N 个泰波那契数
https://leetcode-cn.com/problems/n-th-tribonacci-number/
*/
func tribonacci(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		}
		return 1
	}
	// 定义dp
	dp := make([]int, n+1)
	// 初始化dp
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

/**
爬楼梯
https://leetcode-cn.com/problems/climbing-stairs/
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	// 定义dp
	dp := make([]int, n)
	// 初始化dp
	dp[0] = 1
	dp[1] = 2

	// 遍历dp
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

/**
使用最小花费爬楼梯
https://leetcode-cn.com/problems/min-cost-climbing-stairs/
*/
func minCostClimbingStairs(cost []int) int {
	len := len(cost)
	if len == 2 {
		return Min(cost[0], cost[1])
	}

	// 定义dp
	dp := make([]int, len+1)
	// 初始化dp
	dp[0] = cost[0]
	dp[1] = cost[1]

	// 遍历dp
	for i := 2; i < len; i++ {
		dp[i] = Min(dp[i-1], dp[i-2]) + cost[i]
	}
	dp[len] = Min(dp[len-1], dp[len-2])
	return dp[len]
}

/**
最佳观光组合
https://leetcode-cn.com/problems/best-sightseeing-pair/
*/
func maxScoreSightseeingPair(values []int) int {
	// 定义dp
	max := math.MinInt
	// 初始化值 values[i] + i
	pre := values[0]
	for i := 1; i < len(values); i++ {
		max = Max(max, pre+values[i]-i)
		pre = Max(pre, values[i]+i)
	}
	return max
}

/**
买卖股票的最佳时机
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
*/
func maxProfit(prices []int) int {
	len := len(prices)
	if len < 2 {
		return 0
	}
	/**
	定义dp
	dp[i][0] 用于表示当天不持股的资金
	dp[i][1] 用于表示当天持股的资金
	*/
	dp := make([][]int, len)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// 初始化dp
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	// 遍历dp
	for i := 1; i < len; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}
	return dp[len-1][0]
}

/**
贪心
*/
func maxProfit1(prices []int) int {
	min := math.MaxInt
	res := 0
	for i := 0; i < len(prices); i++ {
		min = Min(min, prices[i])
		res = Max(res, prices[i]-min)
	}
	return res
}
