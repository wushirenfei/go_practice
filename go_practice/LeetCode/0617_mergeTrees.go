// :date 2022/5/1

// https://leetcode-cn.com/problems/merge-two-binary-trees/

package LeetCode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	var tmp, subR1Left, subR2Left, subR1Right, subR2Right *TreeNode
	tmp = root1
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
		subR1Left, subR2Left, subR1Right, subR2Right = root1.Left, root2.Left, root1.Right, root2.Right
	} else {
		if root2 != nil {
			tmp = root2
			subR2Left, subR2Right = root2.Left, root2.Right
		} else {
			subR1Left, subR2Right = root1.Left, root1.Right
		}
	}

	tmp.Left = mergeTrees(subR1Left, subR2Left)
	tmp.Right = mergeTrees(subR1Right, subR2Right)
	return tmp
}