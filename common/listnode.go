package common

// ListNode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildList 从切片构建链表，返回头节点。
func BuildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ToSlice 将链表转换为切片，便于断言。
func ToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
