# exercises-algorithm-go

算法与数据结构练习项目，Go 语言版本。与 [exercises-algorithm](../exercises-algorithm)（Java 版）结构对应。

## 项目结构

```
exercises-algorithm-go/
├── go.mod                  ← 单模块，Go 1.23
├── common/                 ← 公共数据结构（TreeNode、ListNode）
│   ├── treenode.go
│   └── listnode.go
├── leetcode/               ← LeetCode 题目
│   └── p0001/              ← 题号补零到 4 位
│       ├── solution.go     ← 主要解法
│       └── solution_test.go← 测试 + 示例（替代 Java 的 Main.java）
└── ctci/                   ← 《Cracking the Coding Interview》面试题
    └── c0101/
```

## 快速开始

```bash
# 运行所有测试
go test ./...

# 运行单题测试（含示例输出）
go test -v ./leetcode/p0001/

# 运行单题 Benchmark
go test -bench=. ./leetcode/p0001/
```

## 约定

| 内容 | Java 版 | Go 版 |
|------|---------|-------|
| 题目包 | `manfred.exercises.leetcode.p0001` | `package p0001` |
| 主解法 | `Solution.java` | `solution.go` |
| 备选解法 | `Solution2.java` | `solution2.go` |
| 示例/手动测试 | `Main.java` | `Example()` in `solution_test.go` |
| 单元测试 | `SolutionTest.java` | `solution_test.go` |
| 公共结构 | 每题各自定义 | `common/treenode.go` 等 |

## 添加新题目

1. 创建目录：`mkdir leetcode/p0055`
2. 新建 `solution.go`，包名 `package p0055`
3. 新建 `solution_test.go`，写测试用例
4. 运行 `go test ./leetcode/p0055/` 验证

## 相关项目

- [exercises-algorithm](../exercises-algorithm) — Java 版
