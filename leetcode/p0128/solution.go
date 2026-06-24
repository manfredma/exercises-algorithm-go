// Package p0128 — LeetCode 第 128 题「最长连续序列」
//
// 给定一个未排序的整数数组 nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 链接：https://leetcode.cn/problems/longest-consecutive-sequence/
package p0128

func longestConsecutive(nums []int) int {
	reslut := 0
	longestConsecutive := make(map[int]int)
	existed := make(map[int]bool)
	for _, v := range nums {
		if existed[v] {
			continue
		}
		existed[v] = true
		l := longestConsecutive[v-1]
		r := longestConsecutive[v+1]
		longestConsecutive[v-l] = l + r + 1
		longestConsecutive[v+r] = l + r + 1
		reslut = max(reslut, l+r+1)
	}
	return reslut
}
