package offer2

import "testing"

func TestCuttingRope(t *testing.T) {
	print(cuttingRope(120))
}

func TestMyPow(t *testing.T) {
	myPow(2.00000, -2)
}

func TestValidateStackSequences(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 3, 2, 1}
	validateStackSequences(arr1, arr2)
}
