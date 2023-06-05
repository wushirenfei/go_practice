// :date 2022/2/10

package AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, current *ListNode
	l1c, l2c, preV := 0, 0, 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			l1c = l1.Val
			l1 = l1.Next
		} else {
			l1c = 0
		}
		if l2 != nil {
			l2c = l2.Val
			l2 = l2.Next
		} else {
			l2c = 0
		}
		if result == nil {
			cv := l1c + l2c
			if cv >= 10 {
				cv -= 10
				preV = 1
			}
			current = &ListNode{
				Val: cv, Next: nil,
			}
			result = current
		} else {
			cv := l1c + l2c + preV
			if cv >= 10 {
				cv -= 10
				preV = 1
			} else {
				preV = 0
			}
			tmp := &ListNode{
				Val: cv, Next: nil,
			}
			current.Next = tmp
			current = tmp
		}
	}
	if preV != 0 {
		current.Next = &ListNode{Val: preV, Next: nil}
	}
	return result
}
