// :date 2022/4/21


// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

package LeetCode

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	var row []int
	next := []*TreeNode{root}
	for len(next) != 0 {
		row, next = doLevelOrder(next)
		result = append(result, row)
	}
	return result
}

func doLevelOrder(source []*TreeNode) ([]int, []*TreeNode) {
	row := make([]int, len(source))
	next := make([]*TreeNode, 0, 2*len(source))
	for i, node := range source {
		row[i] = node.Val
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
	}

	return row, next
}