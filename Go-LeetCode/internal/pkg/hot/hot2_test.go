package main

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	var strArr = [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	byteArr := make([][]byte, 3)
	for i, _ := range byteArr {
		byteArr[i] = make([]byte, 4)
		for j, _ := range byteArr[i] {
			byteArr[i][j] = strArr[i][j][0]
		}
	}
	str := "ABCCED"
	exist(byteArr, str)
}

func TestFlatten(t *testing.T) {
	root := &TreeNode{0, nil, nil}
	//root.Left = &TreeNode{2, nil, nil}
	//root.Left.Left = &TreeNode{3, nil, nil}
	//root.Left.Right = &TreeNode{4, nil, nil}
	//root.Right = &TreeNode{5, nil, nil}
	//root.Right.Right = &TreeNode{6, nil, nil}
	flatten(root)
	fmt.Println(root)
}

func TestLongestConsecutive(t *testing.T) {
	arr := []int{100, 4, 200, 1, 3, 2}
	longestConsecutive(arr)
}
