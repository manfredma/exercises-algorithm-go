# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# 运行所有测试（带竞态检测）
go test -race ./...

# 运行单题测试并打印示例输出
go test -v ./leetcode/p0001/

# 运行单题 Benchmark
go test -bench=. ./leetcode/p0001/

# 查看测试覆盖率
go test -cover ./...

# 格式化代码
gofmt -w .
goimports -w .
```

## 项目结构与约定

本项目是 LeetCode / CTCI 算法题的 Go 练习库，与同目录的 Java 版 `exercises-algorithm` 结构对应。

```
exercises-algorithm-go/
├── common/          # 公共数据结构：TreeNode、ListNode 及其构建/转换工具函数
├── leetcode/pXXXX/  # 题号补零到 4 位，每题一个独立包
└── ctci/cXXXX/      # 《Cracking the Coding Interview》题目，同样一题一包
```

**每道题的文件约定：**

| 文件 | 说明 |
|------|------|
| `solution.go` | 主解法，包名与目录一致（如 `package p0001`） |
| `solution2.go` | 备选解法（如有） |
| `solution_test.go` | 表驱动单元测试 + `Example()` 函数（替代 Java 的 Main.java） |

**`common/` 提供的工具：**
- `TreeNode` + `BuildTree(vals []int)` — 按层序 BFS 数组（`-1` 表示空节点）构建二叉树
- `ListNode` + `BuildList(vals []int)` / `ToSlice(head)` — 链表构建与转换

## 添加新题目

```bash
mkdir leetcode/p0055
# 新建 solution.go（package p0055）和 solution_test.go
go test ./leetcode/p0055/
```

测试文件中用 `Example()` 函数展示示例输出，`go test -v` 时会打印并验证 `// Output:` 注释。
