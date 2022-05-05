package dp

import (
	"fmt"
	"testing"
)

func TestGetRow(t *testing.T) {
	ans := getRow(3)
	fmt.Println(ans)
}

func TestNumDecodings(t *testing.T) {
	numDecodings("10")
}

func TestMinFallingPathSum(t *testing.T) {
	arr := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	minFallingPathSum(arr)
}
