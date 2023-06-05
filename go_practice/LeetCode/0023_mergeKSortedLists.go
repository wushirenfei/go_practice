// :date 2022/2/12

package LeetCode

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}

	return mergeLists(lists)
}

func mergeLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}
	
	if length == 2 {
		var result, current *ListNode
		l1, l2 := lists[0], lists[1]
		for l1 != nil && l2 != nil {
			var tmp *ListNode
			if l1.Val < l2.Val {
				tmp = l1
				l1 = l1.Next
			} else {
				tmp = l2
				l2 = l2.Next
			}
			if current == nil {
				current = tmp
				result = current
			} else {
				current.Next = tmp
				current = tmp
			}
		}
		if l1 != nil {
			current.Next = l1
		} else {
			current.Next = l2
		}
		return result
	}

	half := length / 2
	return mergeLists([]*ListNode{mergeKLists(lists[0:half]), mergeKLists(lists[half:])})
}

func mergeKListsTimeout(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}

	var head, current *ListNode
	for length > 1 {
		var tmpIndx int
		tmp := 10001 // -10^4 <= lists[i][j] <= 10^4
		for i, l := range lists {
			if l != nil && l.Val < tmp {
				tmp, tmpIndx = l.Val, i
			}
		}
		if current == nil {
			current = lists[tmpIndx]
			head = current
		}
		current.Next = lists[tmpIndx]
		current = current.Next
		lists[tmpIndx] = current.Next
		if lists[tmpIndx] == nil {
			length--
		}
	}

	for _, l := range lists {
		if l != nil {
			current.Next = l
		}
	}

	return head
}