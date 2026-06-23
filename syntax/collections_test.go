// 集合操作：切片、Map、排序、字符串、strconv
package syntax

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ─── 切片（Slice）─────────────────────────────────────────────────────────

// ExampleSlice 切片创建的三种方式
func ExampleSlice() {
	s1 := []int{1, 2, 3} // 字面量

	s2 := make([]int, 3) // len=3, cap=3，零值初始化

	s3 := make([]int, 0, 5)  // len=0, cap=5（预分配，append 密集时用）
	s3 = append(s3, 1, 2, 3) // 追加多个元素
	s3 = append(s3, s1...)   // 展开切片追加

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	// Output:
	// [1 2 3]
	// [0 0 0]
	// [1 2 3 1 2 3]
}

// ExampleSliceOps 切片常用技巧
func ExampleSliceOps() {
	s := []int{1, 2, 3, 4, 5}

	// 子切片（左闭右开，与原底层数组共享内存）
	fmt.Println(s[1:3]) // [2 3]
	fmt.Println(s[:2])  // [1 2]
	fmt.Println(s[3:])  // [4 5]

	// 删除索引 2 的元素（保序）
	del := append([]int{}, s[:2]...)
	del = append(del, s[3:]...)
	fmt.Println(del)

	// 反转
	rev := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	fmt.Println(rev)

	// 深拷贝（独立底层数组，修改互不影响）
	cp := make([]int, len(s))
	copy(cp, s)
	cp[0] = 99
	fmt.Println(s[0], cp[0])
	// Output:
	// [2 3]
	// [1 2]
	// [4 5]
	// [1 2 4 5]
	// [5 4 3 2 1]
	// 1 99
}

// ExampleSlice2D 二维切片（不能用 make([][]int, r, c) 一步到位）
func ExampleSlice2D() {
	rows, cols := 3, 4
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols) // 每行单独分配
	}
	grid[1][2] = 99

	fmt.Println(grid[0])
	fmt.Println(grid[1])
	// Output:
	// [0 0 0 0]
	// [0 0 99 0]
}

// ─── Map ──────────────────────────────────────────────────────────────────

// ExampleMap Map 的创建与常规操作
func ExampleMap() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	fmt.Println(m["apple"])   // 存在：返回值
	fmt.Println(m["cherry"])  // 不存在：返回零值 0，不会 panic

	// 两值赋值：区分"不存在"和"值为零值"
	if v, ok := m["banana"]; ok {
		fmt.Println("banana:", v)
	}

	delete(m, "apple") // 删除（key 不存在时无害）
	fmt.Println(len(m))
	// Output:
	// 5
	// 0
	// banana: 3
	// 1
}

// ExampleMapCount 频率统计（Map 经典用法）
func ExampleMapCount() {
	words := []string{"go", "is", "fun", "go", "is", "go"}
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++ // key 不存在时自动以零值初始化再 +1
	}

	// map 遍历顺序随机，需排序后输出
	keys := make([]string, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, freq[k])
	}
	// Output:
	// fun:1
	// go:3
	// is:2
}

// ExampleMapSet Map 模拟集合（Set）—— 用空结构体节省内存
func ExampleMapSet() {
	set := make(map[int]struct{})
	for _, v := range []int{3, 1, 4, 1, 5, 3} {
		set[v] = struct{}{}
	}
	fmt.Println(len(set)) // 去重后 4 个

	_, exists := set[4]
	fmt.Println(exists)

	_, exists = set[9]
	fmt.Println(exists)
	// Output:
	// 4
	// true
	// false
}

// ─── 排序 ─────────────────────────────────────────────────────────────────

// ExampleSort 内置排序与自定义排序
func ExampleSort() {
	// 基础类型排序（升序）
	nums := []int{5, 2, 8, 1, 9}
	sort.Ints(nums)
	fmt.Println(nums)

	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println(strs)

	// 自定义排序：按绝对值升序
	data := []int{-3, 1, -5, 2}
	sort.Slice(data, func(i, j int) bool {
		ai, aj := data[i], data[j]
		if ai < 0 {
			ai = -ai
		}
		if aj < 0 {
			aj = -aj
		}
		return ai < aj
	})
	fmt.Println(data)

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	// Output:
	// [1 2 5 8 9]
	// [apple banana cherry]
	// [1 2 -3 -5]
	// [9 8 5 2 1]
}

// ExampleBinarySearch sort.Search 二分查找（在有序切片上）
func ExampleBinarySearch() {
	nums := []int{1, 3, 5, 7, 9, 11}

	// SearchInts：找第一个 >= target 的索引
	idx := sort.SearchInts(nums, 6)
	fmt.Println(idx, nums[idx]) // 3 7

	// 通用 sort.Search：f(i) 单调 false→true，返回最小 true 的 i
	pos := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= 7
	})
	fmt.Println(pos, nums[pos])
	// Output:
	// 3 7
	// 3 7
}

// ─── 字符串 ───────────────────────────────────────────────────────────────

// ExampleString 字符串基础：字节 vs rune、range 遍历
func ExampleString() {
	s := "Hello, 世界"
	fmt.Println(len(s))         // 字节数：ASCII 7 + 中文 6 = 13
	fmt.Println(len([]rune(s))) // 字符数（Unicode 码点）：9

	// range 按 rune（Unicode 码点）遍历，i 是字节偏移
	for i, r := range s {
		if i > 2 {
			break
		}
		fmt.Printf("s[%d]=%c\n", i, r)
	}

	// 字符串 → []byte → 修改 → 字符串（字符串本身不可变）
	b := []byte(s)
	b[0] = 'h'
	fmt.Println(string(b[:5]))
	// Output:
	// 13
	// 9
	// s[0]=H
	// s[1]=e
	// s[2]=l
	// hello
}

// ExampleStringOps strings 包常用操作
func ExampleStringOps() {
	s := "Go is awesome"

	fmt.Println(strings.Contains(s, "awesome"))
	fmt.Println(strings.HasPrefix(s, "Go"))
	fmt.Println(strings.HasSuffix(s, "some"))
	fmt.Println(strings.Index(s, "is")) // 返回第一次出现的字节偏移

	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.TrimSpace("  hi  "))
	fmt.Println(strings.Replace(s, "awesome", "fast", 1))

	parts := strings.Split("a,b,c", ",")
	fmt.Println(parts)
	fmt.Println(strings.Join(parts, " | "))
	// Output:
	// true
	// true
	// true
	// 3
	// HELLO
	// hi
	// Go is fast
	// [a b c]
	// a | b | c
}

// ExampleStringBuilder strings.Builder 高效字符串拼接
func ExampleStringBuilder() {
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		sb.WriteByte(byte('a' + i)) // 逐字节写入
	}
	fmt.Println(sb.String())

	// WriteString 批量写入
	sb.Reset()
	for i, w := range []string{"hello", "world"} {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(w)
	}
	fmt.Println(sb.String())
	// Output:
	// abcde
	// hello world
}

// ExampleStrconv 数字与字符串互转
func ExampleStrconv() {
	// int → string
	s := strconv.Itoa(42)
	fmt.Println(s)

	// string → int（返回 error，需处理）
	n, err := strconv.Atoi("123")
	fmt.Println(n, err)

	// 非法输入
	_, err = strconv.Atoi("abc")
	fmt.Println(err != nil)

	// 其他进制
	fmt.Println(strconv.FormatInt(255, 2))  // 十进制 → 二进制字符串
	fmt.Println(strconv.FormatInt(255, 16)) // 十进制 → 十六进制字符串
	// Output:
	// 42
	// 123 <nil>
	// true
	// 11111111
	// ff
}
