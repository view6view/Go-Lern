package str

/**
实现 strStr()
https://leetcode.cn/problems/implement-strstr/
*/
func strStr(haystack string, needle string) int {
	first := needle[0]
	valueCount := len(haystack)
	strCount := len(needle)
	max := valueCount - strCount
	for i := 0; i <= max; i++ {
		if haystack[i] == first {
			if haystack[i:i+strCount] == needle {
				return i
			}
		}
	}
	return -1
}

/**
移除元素
https://leetcode.cn/problems/remove-element/
*/
func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
