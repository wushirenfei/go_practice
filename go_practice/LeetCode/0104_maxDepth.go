// :date 2022/4/21

package LeetCode

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	if maxLeft > maxRight {
		return maxLeft + 1
	}
	return maxRight + 1
}
