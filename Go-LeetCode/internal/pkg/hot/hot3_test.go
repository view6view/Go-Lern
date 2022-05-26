package main

import "testing"

func TestRob(t *testing.T) {
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Left.Right = &TreeNode{3, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{1, nil, nil}
	rob(root)
}

func TestDecodeString(t *testing.T) {
	decodeString("100[leetcode]")
}
