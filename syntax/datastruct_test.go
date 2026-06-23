// 算法常用数据结构：结构体、栈、队列、堆（优先队列）
package syntax

import (
	"container/heap"
	"fmt"
	"sort"
)

// ─── 结构体与方法 ──────────────────────────────────────────────────────────

// Interval 区间（典型算法结构体）
type Interval struct {
	Start, End int
}

// ExampleStruct 结构体、匿名结构体、结构体切片排序
func ExampleStruct() {
	// 具名结构体
	iv := Interval{Start: 1, End: 5}
	fmt.Println(iv.Start, iv.End)

	// 匿名结构体（一次性使用，无需命名）
	point := struct{ X, Y int }{X: 3, Y: 4}
	fmt.Println(point.X, point.Y)

	// 结构体切片 + 自定义排序（区间按起点升序）
	intervals := []Interval{{3, 6}, {1, 4}, {2, 5}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	for i, v := range intervals {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("[%d,%d]", v.Start, v.End)
	}
	fmt.Println()
	// Output:
	// 1 5
	// 3 4
	// [1,4] [2,5] [3,6]
}

// ─── 栈（Stack）──────────────────────────────────────────────────────────

// ExampleStack 用切片模拟 LIFO 栈
func ExampleStack() {
	var stack []int

	// Push
	stack = append(stack, 1, 2, 3)
	fmt.Println("after push:", stack)

	// Top（不弹出）
	top := stack[len(stack)-1]
	fmt.Println("top:", top)

	// Pop
	stack = stack[:len(stack)-1]
	fmt.Println("after pop:", stack)

	fmt.Println("empty:", len(stack) == 0)
	// Output:
	// after push: [1 2 3]
	// top: 3
	// after pop: [1 2]
	// empty: false
}

// ─── 队列（Queue）────────────────────────────────────────────────────────

// ExampleQueue 用切片模拟 FIFO 队列（BFS 常用）
func ExampleQueue() {
	var queue []int

	// Enqueue
	queue = append(queue, 10, 20, 30)
	fmt.Println("after enqueue:", queue)

	// Dequeue
	front := queue[0]
	queue = queue[1:]
	fmt.Println("dequeued:", front)
	fmt.Println("remaining:", queue)
	// Output:
	// after enqueue: [10 20 30]
	// dequeued: 10
	// remaining: [20 30]
}

// ExampleDeque 用切片模拟双端队列（单调队列等场景）
func ExampleDeque() {
	var dq []int

	// 两端 push
	dq = append(dq, 1, 2, 3)        // push back
	dq = append([]int{0}, dq...)    // push front

	// 两端 pop
	front, back := dq[0], dq[len(dq)-1]
	dq = dq[1 : len(dq)-1]

	fmt.Println("front:", front)
	fmt.Println("back:", back)
	fmt.Println("middle:", dq)
	// Output:
	// front: 0
	// back: 3
	// middle: [1 2]
}

// ─── 堆（Heap / 优先队列）container/heap ──────────────────────────────────
//
// 实现 heap.Interface（5 个方法）即可获得 O(log n) 的 Push/Pop。
// 堆顶（index 0）是 Less(i,j) 关系中"最小"的元素。

// IntMinHeap 整数最小堆
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆：较小值优先
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// ExampleIntMinHeap 最小堆
func ExampleIntMinHeap() {
	h := &IntMinHeap{5, 2, 8, 1}
	heap.Init(h) // O(n) 建堆

	heap.Push(h, 3)
	fmt.Println("min:", (*h)[0]) // 堆顶始终是最小值

	for h.Len() > 0 {
		fmt.Print(heap.Pop(h)) // 每次弹出当前最小值
		if h.Len() > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	// Output:
	// min: 1
	// 1 2 3 5 8
}

// IntMaxHeap 整数最大堆（仅 Less 取反）
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 最大堆：较大值优先
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntMaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// ExampleIntMaxHeap 最大堆
func ExampleIntMaxHeap() {
	h := &IntMaxHeap{5, 2, 8, 1}
	heap.Init(h)
	fmt.Println("max:", (*h)[0])

	for h.Len() > 0 {
		fmt.Print(heap.Pop(h))
		if h.Len() > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	// Output:
	// max: 8
	// 8 5 2 1
}

// ─── 自定义结构体堆（优先队列经典写法）────────────────────────────────────

// Task 带优先级的任务（算法题中常见的 pair 结构）
type Task struct {
	Name     string
	Priority int
}

// TaskHeap 按优先级排列的最小堆（priority 越小越先处理）
type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x any)        { *h = append(*h, x.(Task)) }
func (h *TaskHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// ExampleTaskHeap 自定义结构体优先队列
func ExampleTaskHeap() {
	h := &TaskHeap{}
	heap.Init(h)

	heap.Push(h, Task{"C", 5})
	heap.Push(h, Task{"A", 1})
	heap.Push(h, Task{"B", 3})

	for h.Len() > 0 {
		t := heap.Pop(h).(Task)
		fmt.Printf("priority=%d name=%s\n", t.Priority, t.Name)
	}
	// Output:
	// priority=1 name=A
	// priority=3 name=B
	// priority=5 name=C
}
