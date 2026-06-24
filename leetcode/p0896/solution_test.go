package p0896

import (
	"fmt"
	"testing"
)

func TestIsMonotonic(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 2, 3}, true},
		{[]int{6, 5, 4, 4}, true},
		{[]int{1, 3, 2}, false},
		{[]int{1}, true},
		{[]int{1, 1, 1}, true},
	}
	for _, tc := range tests {
		got := isMonotonic(tc.nums)
		if got != tc.want {
			t.Errorf("isMonotonic(%v) = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

// Example 演示用法，go test -v 时会打印输出。
func Example() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
	// Output:
	// true
	// true
	// false
}
