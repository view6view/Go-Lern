package dp

import "math"

/**
跳跃游戏 II
https://leetcode-cn.com/problems/jump-game-ii/
*/
func jump(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	dp[0] = 0
	for i := 1; i < len; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len; i++ {
		canJump := Min(len-1, i+nums[i])
		for j := i + 1; j <= canJump; j++ {
			dp[j] = Min(dp[j], dp[i]+1)
		}
	}
	return dp[len-1]
}

/**
分割等和子集
https://leetcode-cn.com/problems/partition-equal-subset-sum/
*/
func canPartition(nums []int) bool {
	len := len(nums)
	if len == 1 {
		return false
	}
	sum := nums[0]
	for i := 1; i < len; i++ {
		sum += nums[i]
	}

	// 总和为奇数，不能平分
	if sum&1 == 1 {
		return false
	}
	avg := sum / 2
	dp := make([]int, avg+1)
	// 正向遍历物品
	for _, num := range nums {
		// 逆向遍历背包
		for i := avg; i >= num; i-- {
			// 更新背包能放的容积
			dp[i] = Max(dp[i], dp[i-num]+num)
		}
	}
	return dp[avg] == avg
}

/**
最后一块石头的重量 II
https://leetcode-cn.com/problems/last-stone-weight-ii/
*/
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	avg := sum >> 1

	dp := make([]int, avg+1)

	// 正向遍历物品
	for _, stone := range stones {
		// 逆向遍历背包
		for i := avg; i >= stone; i-- {
			dp[i] = Max(dp[i], dp[i-stone]+stone)
		}
	}
	return sum - 2*dp[avg]
}

/**
目标和
https://leetcode-cn.com/problems/target-sum/
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	target += sum

	if target&1 == 1 {
		return 0
	}

	target = Abs(target / 2)

	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

/**
一和零
https://leetcode-cn.com/problems/ones-and-zeroes/
*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for strIdx, _ := range strs {
		zero := 0
		one := 0
		for charIdx, _ := range strs[strIdx] {
			if strs[strIdx][charIdx] == '0' {
				zero++
			} else {
				one++
			}
		}

		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = Max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

/**
零钱兑换 II
https://leetcode-cn.com/problems/coin-change-2/
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}

/**
组合总和 Ⅳ
https://leetcode-cn.com/problems/combination-sum-iv/
*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
