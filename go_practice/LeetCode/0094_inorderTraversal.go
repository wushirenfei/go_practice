// :date 2022/4/18

package LeetCode

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	return doInorderTraversal(root, result)
}

func doInorderTraversal(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = doInorderTraversal(root.Left, result)
	result = append(result, root.Val)
	result = doInorderTraversal(root.Right, result)
	return result
}