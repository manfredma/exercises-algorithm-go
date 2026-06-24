package p0128

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
		{[]int{}, 0},
		{[]int{1}, 1},
	}
	for _, tc := range tests {
		got := longestConsecutive(tc.nums)
		if got != tc.want {
			t.Errorf("longestConsecutive(%v) = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

// Example 演示用法，go test -v 时会打印输出。
func Example() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	// Output:
	// 4
	// 9
}
