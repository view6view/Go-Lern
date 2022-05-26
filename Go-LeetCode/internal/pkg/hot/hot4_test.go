package main

import "testing"

func TestReconstructQueue(t *testing.T) {
	arr := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue(arr)
}

func TestFindDisappearedNumbers(t *testing.T) {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	findDisappearedNumbers(arr)
}

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	//root := strToTreeNode("[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]")
	diameterOfBinaryTree(root)
}
