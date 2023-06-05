// :date 2022/4/30

package LeetCode

// https://leetcode-cn.com/problems/invert-binary-tree/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}