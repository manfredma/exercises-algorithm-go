// Package p0896 — LeetCode 第 896 题「单调数列」
//
// 如果数组是单调递增或单调递减的，那么它是单调的。
// 当给定的数组 nums 是单调数组时返回 true，否则返回 false。
//
// 链接：https://leetcode.cn/problems/monotonic-array/
package p0896

func isMonotonic(nums []int) bool {
	bigger := false
	smaller := false
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			bigger = true
		} else if nums[i] < nums[i-1] {
			smaller = true
		}
	}
	return !(bigger && smaller)
}
