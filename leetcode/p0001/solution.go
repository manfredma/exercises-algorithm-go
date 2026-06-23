// Package p0001 — LeetCode 第 1 题「两数之和」
//
// 给定一个整数数组 nums 和一个目标值 target，在数组中找出和为目标值的两个整数的索引。
// 假设每种输入只有一个答案，且同一个元素不能使用两遍。
//
// 链接：https://leetcode.cn/problems/two-sum/
package p0001

// twoSum 哈希表一次遍历，时间 O(n)，空间 O(n)。
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return nil
}
