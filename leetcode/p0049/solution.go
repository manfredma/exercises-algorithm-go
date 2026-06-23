// Package p0049 — LeetCode 第 49 题「字母异位词分组」
//
// 给你一个字符串数组，请你将字母异位词组合在一起。
// 可以按任意顺序返回结果列表。
//
// 链接：https://leetcode.cn/problems/group-anagrams/
package p0049

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, s := range strs {
		var key [26]int
		for _, ch := range s {
			key[ch-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	var result [][]string
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}
