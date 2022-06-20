package offer2

import (
	"fmt"
	"math"
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

/**
数组中数字出现的次数
https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
*/
func singleNumbers(nums []int) []int {
	x, y, n, m := 0, 0, 0, 1
	for _, num := range nums {
		n ^= num
	}
	// 此时n的值为 x ^ y
	// x 和 y 的值不相同，x ^ y 的某二进制位为1，则x 与 y在此二进制位上一定不同，找出这个位置
	// 或者m = n - (n & (n - 1)) 原理是 n & (n-1)可以把n二进制最右边的1变为0，然后拿n减去这个值，就是我们需要的m了
	for m&n == 0 {
		m <<= 1
	}
	for _, num := range nums {
		// 通过&m将num分为两种情况，则一定可以把x y 区分开来
		if num&m == 0 {
			x ^= num
		} else {
			y ^= num
		}
	}
	return []int{x, y}
}

/**
丑数
https://leetcode.cn/problems/chou-shu-lcof/
*/
func nthUglyNumber(n int) int {
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := 2*dp[a], 3*dp[b], 5*dp[c]
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}

/**
数组中数字出现的次数 II
https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
*/
func singleNumber(nums []int) int {
	a := 0
	b := 0
	for _, c := range nums {
		a, b = a&^c|b&c, b&^c|^a&^b&c
	}
	return ^a & b
}

/**
和为s的连续正数序列
https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
*/
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	i, j, s := 1, 2, 3
	for i < j {
		if s == target {
			nums := make([]int, j-i+1)
			for k := i; k <= j; k++ {
				nums[k-i] = k
			}
			res = append(res, nums)
		}
		if s >= target {
			s -= i
			i++
		} else {
			j++
			s += j
		}
	}
	return res
}

/**
n个骰子的点数
https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/
*/
func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := range dp {
		dp[i] = 1.0 / 6.0
	}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, i*5+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / 6.0
			}
		}
		dp = tmp
	}
	return dp
}

/**
圆圈中最后剩下的数字
https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
*/
func lastRemaining(n int, m int) int {
	f := 0
	for i := 2; i <= n; i++ {
		f = (f + m) % i
	}
	return f
}

/**
股票的最大利润
https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
*/
func maxProfit(prices []int) int {
	var Min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var Max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min, profit := math.MaxInt, 0
	for _, price := range prices {
		min = Min(min, price)
		profit = Max(profit, price-min)
	}
	return profit
}

/**
求1+2+…+n
https://leetcode.cn/problems/qiu-12n-lcof/
*/
func sumNums(n int) int {
	//return n * (n + 1) / 2
	_ = n > 0 && func() bool { n += sumNums(n - 1); return true }()
	return n
}

/**
不用加减乘除做加法
https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
*/
func add(a int, b int) int {
	for b != 0 { // 当进位为 0 时候跳出
		c := (a & b) << 1 // c = 进位
		a ^= b            // a = 非进位和
		b = c             // b = 进位
	}
	return a
}

/**
构建乘积数组
https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/
*/
func constructArr(a []int) []int {
	zeroLastIdx := -1
	zeroCount := 0
	total := 1
	for i, num := range a {
		if num == 0 {
			zeroLastIdx = i
			zeroCount++
		} else {
			total *= num
		}
	}
	if zeroCount == 0 {
		for i, num := range a {
			a[i] = total / num
		}
	} else if zeroCount == 1 {
		for i := range a {
			if i != zeroLastIdx {
				a[i] = 0
			} else {
				a[i] = total
			}
		}
	} else {
		for i := range a {
			a[i] = 0
		}
	}

	return a
}

/**
把字符串转换成整数
https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
*/
func strToInt(str string) int {
	res, bndry := 0, math.MaxInt32/10
	idx, sign, len := 0, 1, len(str)
	if len == 0 {
		return 0
	}
	for str[idx] == ' ' {
		idx++
		if idx == len {
			return 0
		}
	}
	if str[idx] == '-' {
		sign = -1
	}
	if str[idx] == '-' || str[idx] == '+' {
		idx++
	}
	for ; idx < len; idx++ {
		if str[idx] < '0' || str[idx] > '9' {
			break
		}
		if res > bndry || (res == bndry && str[idx] > '7') {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + int(str[idx]-'0')
	}
	return sign * res
}
