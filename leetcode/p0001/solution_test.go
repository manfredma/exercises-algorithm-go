package p0001

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tc := range tests {
		got := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
		}
	}
}

// Example 演示用法，go test -v 时会打印输出（等价于 Java 的 Main.java）。
func Example() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
	// Output:
	// [0 1]
}
