package dp

import "testing"

func TestJump(t *testing.T) {
	arr := []int{2, 3, 1, 1, 4}
	jump(arr)
}

func TestCanPartition(t *testing.T) {
	arr := []int{1, 5, 11, 5}
	canPartition(arr)
}

func TestFindMaxForm(t *testing.T) {
	arr := []string{"10", "0001", "111001", "1", "0"}
	findMaxForm(arr, 5, 3)
}
