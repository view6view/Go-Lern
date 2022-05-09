package main

/**
两数之和
https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	idxMap := map[int]int{}
	for i, num := range nums {
		if v, ok := idxMap[target-num]; ok {
			return []int{v, i}
		}
		idxMap[num] = i
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
两数相加
https://leetcode-cn.com/problems/add-two-numbers/
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	//定义一个尾结点，或者可以理解为临时节点
	var tail *ListNode
	//余数carry
	carry := 0
	//依次遍历两个链表，只要元素不为空就进行下一步
	for l1 != nil || l2 != nil {
		//定义两个变量存储各个节点的值
		n1, n2 := 0, 0
		//从第一个链表开始
		if l1 != nil {
			//把每个节点的值赋给n1
			n1 = l1.Val
			//节点后移
			l1 = l1.Next
		}
		//l2同上
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//此时是两个链表第一个元素的和 + 余数
		sum := n1 + n2 + carry
		//sum%10是节点的当前值，如果是10,取余后当前节点值为0，sum/10是求十位的那个数
		sum, carry = sum%10, sum/10
		//此时申请一个新的链表存储两个链表的和
		if head == nil {
			//申请新的链表
			head = &ListNode{Val: sum}
			//这一步是为了保持头结点不变的情况下指针可以右移，所以说tail相当于临时节点，理解成尾节点也可以，因
			//为此时新链表中只有一个节点，所以头结点和尾结点都指向同一个元素。
			tail = head
		} else {
			//第二个节点后开始逐渐往尾结点增加元素
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	//把最后一位的余数加到链表最后。
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

/**
下一个排列
https://leetcode-cn.com/problems/next-permutation/
*/
func nextPermutation(nums []int) {
	len := len(nums)

	if len <= 1 {
		return
	}

	i, j, k := len-2, len-1, len-1

	// 找到num[i] < nums[j]
	for j > 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if j > 0 {
		// 从最后一个数向前开始找到第一个大于 nums[i]的数据
		for nums[i] >= nums[k] {
			k--
		}
		// 交换nums[i] nums[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 反转nums[j:end]
	for begin, end := j, len-1; begin < end; begin, end = begin+1, end-1 {
		// 交换
		nums[begin], nums[end] = nums[end], nums[begin]
	}
}

/**
https://leetcode.cn/problems/search-in-rotated-sorted-array/
搜索旋转排序数组
*/
func search(nums []int, target int) int {
	len := len(nums)
	if len == 0 {
		return -1
	}
	left, right := 0, len-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] < nums[right] { // 右边有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // 左边有序
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

/**
在排序数组中查找元素的第一个和最后一个位置
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
*/
func searchRange(nums []int, target int) []int {
	len := len(nums)
	left, right := 0, len-1
	if len == 0 || target < nums[left] || target > nums[right] {
		return []int{-1, -1}
	}
	for left <= right {
		mid := (left + right) / 2

		// 找到值
		if target == nums[mid] {
			left = mid - 1
			right = mid + 1
			for left > -1 && nums[left] == nums[mid] {
				left--
			}
			for right < len && nums[right] == nums[mid] {
				right++
			}
			return []int{left + 1, right - 1}
		}

		if target < nums[mid] { // 左边
			right = mid - 1
			for right >= left && nums[right] == nums[mid] {
				right--
			}
		} else { // 右边
			left = mid + 1
			for left <= right && nums[left] == nums[mid] {
				left++
			}
		}
	}
	return []int{-1, -1}
}
