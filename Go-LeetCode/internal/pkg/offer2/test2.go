package offer2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
复杂链表的复制
https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	node := head

	m := map[*Node]int{}
	idx := 0
	nodes := make([]*Node, 0)
	for node != nil {
		m[node] = idx
		nodes = append(nodes, &Node{Val: node.Val})
		idx++
		node = node.Next
	}

	idx = 0
	node = head
	for node.Next != nil {
		if node.Random != nil {
			nodes[idx].Random = nodes[m[node.Random]]
		}
		nodes[idx].Next = nodes[idx+1]
		idx++
		node = node.Next
	}
	if node.Random != nil {
		nodes[idx].Random = nodes[m[node.Random]]
	}
	return nodes[0]
}

/**
字符串的排列
https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/
*/
func permutation(s string) []string {
	bytes := []byte(s)
	end := len(bytes) - 1
	res := make([]string, 0)

	type void struct{}

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == end {
			res = append(res, string(bytes))
		}
		set := map[byte]void{}
		for i := idx; i <= end; i++ {
			if _, ok := set[bytes[i]]; !ok {
				set[bytes[i]] = void{}
				bytes[i], bytes[idx] = bytes[idx], bytes[i]
				dfs(idx + 1)
				bytes[i], bytes[idx] = bytes[idx], bytes[i]
			}
		}
	}
	dfs(0)
	return res
}

/**
把数组排成最小的数
https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
字符串 xy < yx , yz < zy ，需证明 xz < zx 一定成立。

设十进制数 x, y, z 分别有 a, b, c 位，则有：
（左边是字符串拼接，右边是十进制数计算，两者等价）
xy = x * 10^b + y
yx = y * 10^a + x

则 xy < yx 可转化为：
x * 10^b + y < y * 10^a + x
x (10^b - 1) < y (10^a - 1)
x / (10^a - 1) < y / (10^b - 1)     ①

同理， 可将 yz < zy 转化为：
y / (10^b - 1) < z / (10^c - 1)     ②

将 ① ② 合并，整理得：
x / (10^a - 1) < y / (10^b - 1) < z / (10^c - 1)
x / (10^a - 1) < z / (10^c - 1)
x (10^c - 1) < z (10^a - 1)
x * 10^c + z < z * 10^a + x
∴  可推出 xz < zx ，传递性证毕
*/
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})
	builder := strings.Builder{}
	for _, num := range nums {
		builder.WriteString(fmt.Sprintf("%d", num))
	}
	return builder.String()
}

/**
数字序列中某一位的数字
https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
*/
func findNthDigit(n int) int {
	digit := 1
	start := 1
	count := 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9
	}
	num := start + (n-1)/digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}

/**
把数字翻译成字符串
https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
*/
func translateNum(num int) int {
	str := strconv.Itoa(num)
	len := len(str)
	dp := make([]int, len+1)
	dp[0] = 1
	dp[1] = 1
	// 用于标记当前遍历的char，它的前2个char是否可以用来翻译
	// 如果true，说明当前存在间隔，dp[i] = dp[i-1] * 2
	// 如果false，说明前2个char可以用来翻译并且已经用过了，dp[i] = dp[i-1] + dp[i-2]
	isolation := false
	for i := 2; i <= len; i++ {
		cur, _ := strconv.Atoi(str[i-2 : i])
		if cur < 26 && cur > 9 {
			if isolation {
				dp[i] = dp[i-1] * 2
				isolation = false
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
			isolation = true
		}
	}
	return dp[len]
}

/**
礼物的最大价值
https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/
*/
func maxValue(grid [][]int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[m][n]
}

/**
最长不含重复字符的子字符串
https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
*/
func lengthOfLongestSubstring(s string) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 由于go map的特性，记录的下标比真实下标+1
	m := map[byte]int{}
	len := len(s)
	res := 0
	pre := 0
	for i := 0; i < len; i++ {
		preIdx := m[s[i]]
		if preIdx != 0 {
			pre = max(pre, preIdx)
		}
		m[s[i]] = i + 1
		res = max(res, i-pre+1)
	}
	return res
}
