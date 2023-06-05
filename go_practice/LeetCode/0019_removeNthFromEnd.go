// :date 2022/2/12

package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Note:
// 1. When lenLinkedList == n == 1，return nil
// 2. When lenLinkedList == n && n > 1, return head.Next
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current, nSlice := head, make([]*ListNode, 0, n+1)
	for current.Next != nil {
		nSlice = append(nSlice, current)
		if len(nSlice) == n+1 {
			nSlice = nSlice[1:]
		}
		current = current.Next
	}
	if len(nSlice) == n-1 {
		if n == 1 {
			return nil
		} else {
			return head.Next
		}
	}
	if n == 1 {
		nSlice[0].Next = nil
	} else {
		nSlice[0].Next = nSlice[1].Next
	}
	return head
}
