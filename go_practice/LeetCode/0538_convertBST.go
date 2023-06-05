// :date 2022/5/3

package LeetCode

// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

func convertBST(root *TreeNode) *TreeNode {
	InOrderRight(root, 0)
	return root
}

func InOrderRight(root *TreeNode, nowSum int) int {
	if root == nil {
		return nowSum
	}
	nowSum = InOrderRight(root.Right, nowSum)
	nowSum += root.Val
	root.Val = nowSum
	nowSum = InOrderRight(root.Left, nowSum)
	return nowSum
}