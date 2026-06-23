package p0049

import (
	"fmt"
	"sort"
	"testing"
)

// normalize 对结果做规范化排序，使比较与顺序无关。
// 每个子数组内部排序，再按首元素对子数组排序。
func normalize(groups [][]string) [][]string {
	for _, g := range groups {
		sort.Strings(g)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 {
			return true
		}
		if len(groups[j]) == 0 {
			return false
		}
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func equalGroups(a, b [][]string) bool {
	a = normalize(a)
	b = normalize(b)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
	}
	for _, tc := range tests {
		got := groupAnagrams(tc.strs)
		if !equalGroups(got, tc.want) {
			t.Errorf("groupAnagrams(%v) = %v, want %v", tc.strs, got, tc.want)
		}
	}
}

// Example 演示用法，go test -v 时会打印输出。
func Example() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	// 规范化后打印，确保输出稳定
	normalized := normalize(result)
	fmt.Println(normalized)
	// Output:
	// [[bat] [nat tan] [ate eat tea]]
}
