// :date 2022/4/24

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

package LeetCode

func flatten(root *TreeNode)  {
	tmpSlice := preOrder(root, nil)
	if tmpSlice == nil || len(tmpSlice) ==0 {
		return
	}

	for i := 0; i < len(tmpSlice) - 1; i++ {
		tmp := tmpSlice[i]
		tmp.Left, tmp.Right = nil, tmpSlice[i+1]
	}
}

func preOrder(root *TreeNode, result []*TreeNode) []*TreeNode {
	if root == nil {
		return result
	}
	result = append(result, root)
	result = preOrder(root.Left, result)
	result = preOrder(root.Right, result)
	return result
}

