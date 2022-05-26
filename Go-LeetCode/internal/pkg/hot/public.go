package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
类似于把数字字符转换为数型结构
[4,-7,-3,null,null,-9,...]
*/
func strToTreeNode(s string) *TreeNode {
	lS := len(s)
	s = s[1 : lS-1]
	splitArr := strings.Split(s, ",")
	lSpilt := len(splitArr)
	nodes := make([]*TreeNode, lSpilt)
	for i, str := range splitArr {
		if str == "null" {
			nodes[i] = nil
		} else {
			num, _ := strconv.Atoi(str)
			nodes[i] = &TreeNode{Val: num}
		}
	}
	for i := 0; i < (lSpilt-2)/2; i++ {
		nodes[i].Left = nodes[i*2+1]
		nodes[i].Right = nodes[i*2+2]
	}
	return nodes[0]
}
