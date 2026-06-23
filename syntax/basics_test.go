// 语言基础：变量、控制流、函数、闭包、指针、数学、位运算
package syntax

import (
	"fmt"
	"math"
	"strings"
)

// ExampleVariable 变量声明：短变量、var、const、零值
func ExampleVariable() {
	a := 42              // 最常用：短变量声明（编译器推断类型）
	var b int            // 零值 0（显式声明，无初始值）
	var c, d = 1, "hi"  // 多变量声明
	const Pi = 3.14      // 常量（编译期确定）

	fmt.Println(a, b, c, d, Pi)
	// Output:
	// 42 0 1 hi 3.14
}

// ExampleTypeConvert 类型转换（Go 不做隐式转换，必须显式）
func ExampleTypeConvert() {
	i := 65
	f := float64(i)   // int → float64
	r := rune(i)      // int → rune（Unicode 码点）
	s := string(r)    // rune → string（注意：string(65) 是 "A" 而非 "65"）

	fmt.Println(i, f, r, s)
	// Output:
	// 65 65 65 A
}

// ExampleIfElse if/else 与初始化语句
func ExampleIfElse() {
	x := 85
	if x >= 90 {
		fmt.Println("A")
	} else if x >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// if 内部声明变量（作用域仅限 if/else 块）
	if v := x * 2; v > 100 {
		fmt.Println("doubled:", v)
	}
	// Output:
	// B
	// doubled: 170
}

// ExampleSwitch switch 语句（case 支持多值；无需 break）
func ExampleSwitch() {
	for _, n := range []int{1, 3, 7} {
		switch n {
		case 1, 2:
			fmt.Println(n, "→ small")
		case 3, 4, 5:
			fmt.Println(n, "→ medium")
		default:
			fmt.Println(n, "→ large")
		}
	}
	// Output:
	// 1 → small
	// 3 → medium
	// 7 → large
}

// ExampleForLoop for 循环的三种形式 + range
func ExampleForLoop() {
	// 传统 C 风格
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	// range 遍历（索引 + 值）
	for i, v := range []string{"a", "b", "c"} {
		fmt.Printf("%d=%s\n", i, v)
	}

	// while 风格（省略初始化和后置语句）
	n := 1
	for n < 10 {
		n *= 2
	}
	fmt.Println(n)
	// Output:
	// 15
	// 0=a
	// 1=b
	// 2=c
	// 16
}

// ExampleLabelBreak 带标签的 break（跳出嵌套循环，算法题常用）
func ExampleLabelBreak() {
	var result []string
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+j >= 3 {
				break outer // 直接跳出外层循环
			}
			result = append(result, fmt.Sprintf("(%d,%d)", i, j))
		}
	}
	fmt.Println(strings.Join(result, " "))
	// Output:
	// (0,0) (0,1) (0,2) (1,0) (1,1)
}

// ─── 函数 ─────────────────────────────────────────────────────────────────

// minMax 演示：命名返回值 + 裸 return
func minMax(nums []int) (lo, hi int) {
	lo, hi = nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < lo {
			lo = v
		}
		if v > hi {
			hi = v
		}
	}
	return // 裸 return：返回命名变量的当前值
}

// sumAll 演示：可变参数（...int）
func sumAll(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// ExampleFunction 多返回值、命名返回值、可变参数
func ExampleFunction() {
	lo, hi := minMax([]int{3, 1, 4, 1, 5, 9})
	fmt.Println(lo, hi)

	fmt.Println(sumAll(1, 2, 3, 4, 5))

	nums := []int{10, 20, 30}
	fmt.Println(sumAll(nums...)) // 用 ... 将切片展开为可变参数
	// Output:
	// 1 9
	// 15
	// 60
}

// ExampleClosure 闭包：函数捕获外层变量
func ExampleClosure() {
	// adder 返回一个闭包，捕获 base
	adder := func(base int) func(int) int {
		return func(x int) int { return base + x }
	}

	add5 := adder(5)
	fmt.Println(add5(3))
	fmt.Println(add5(10))
	// Output:
	// 8
	// 15
}

// ─── 指针 ─────────────────────────────────────────────────────────────────

func triple(p *int) { *p *= 3 }

// ExamplePointer 指针：传地址修改原值
func ExamplePointer() {
	x := 7
	triple(&x) // 传指针，函数内修改影响原变量
	fmt.Println(x)

	p := new(int) // 分配零值 int，返回 *int
	*p = 42
	fmt.Println(*p)
	// Output:
	// 21
	// 42
}

// ─── 数学 ─────────────────────────────────────────────────────────────────

// ExampleMath 数学操作（Go 1.21+ 内置 min/max，适用于任意有序类型）
func ExampleMath() {
	fmt.Println(min(3, 7))        // 内置，Go 1.21+
	fmt.Println(max(3, 7))        // 内置，Go 1.21+
	fmt.Println(math.Abs(-5.0))   // float64 绝对值
	fmt.Println(int(math.Sqrt(16)))
	fmt.Println(math.MaxInt)      // 平台相关最大值（64 位系统）
	// Output:
	// 3
	// 7
	// 5
	// 4
	// 9223372036854775807
}

// ─── 位运算 ───────────────────────────────────────────────────────────────

// ExampleBitOp 位运算（算法题高频）
func ExampleBitOp() {
	n := 0b1010 // 10（二进制字面量，Go 1.13+）

	fmt.Println(n & 0b1100)  // AND  → 0b1000 = 8
	fmt.Println(n | 0b0101)  // OR   → 0b1111 = 15
	fmt.Println(n ^ 0b1100)  // XOR  → 0b0110 = 6
	fmt.Println(n >> 1)      // 右移  → 0b0101 = 5
	fmt.Println(n << 1)      // 左移  → 0b10100 = 20
	fmt.Println(n & (n - 1)) // 清除最低位 1 → 0b1000 = 8
	fmt.Println(n & (-n))    // 取最低位 1   → 0b0010 = 2
	// Output:
	// 8
	// 15
	// 6
	// 5
	// 20
	// 8
	// 2
}
