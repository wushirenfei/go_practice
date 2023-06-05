// :date 2022/4/18

package LeetCode

import "fmt"

// https://leetcode-cn.com/problems/validate-binary-search-tree/

type BSTNode struct {
	Max int
	Min int
}

var NodeMap = map[*TreeNode]*BSTNode{}

func isValidBST(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		NodeMap[root] = &BSTNode{Max: root.Val, Min: root.Val}
		return true
	}

	isLeft, isRight := true, true
	if root.Left != nil {
		isLeft = isValidBST(root.Left)
	}
	if !isLeft {
		return false
	}
	if root.Right != nil {
		isRight = isValidBST(root.Right)
	}
	if !isRight {
		return false
	}

	left, right := NodeMap[root.Left], NodeMap[root.Right]
	if (root.Left == nil || root.Val > left.Max) && (root.Right == nil || root.Val < right.Min) {
		tmp := &BSTNode{Max: root.Val, Min: root.Val}
		if root.Left != nil {
			tmp.Min = left.Min
		}
		if root.Right != nil {
			tmp.Max = right.Max
		}
		NodeMap[root] = tmp
		return true
		fmt.Println()
	}

	return false
}
