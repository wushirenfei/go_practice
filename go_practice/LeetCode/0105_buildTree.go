// :date 2022/4/22

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package LeetCode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return node
	}

	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	leftInorder, rightInorder := inorder[0:i], inorder[i+1:]
	if len(leftInorder) == 0 {
		node.Right = buildTree(preorder[1:], rightInorder)
		return node
	} else if len(rightInorder) == 0 {
		node.Left = buildTree(preorder[1:], leftInorder)
		return node
	}
	leftElem := make(map[int]bool, len(leftInorder))
	for _, e := range leftInorder {
		leftElem[e] = true
	}
	j := 1
	for ; j < len(preorder); j++ {
		if !leftElem[preorder[j]] {
			break
		}
	}
	leftPreorder, rightPreorder := preorder[1:j], preorder[j:]

	// faster
	//var leftPreorder, rightPreorder []int
	//for j:=1; j<len(preorder); j++ {
	//	if !leftElem[preorder[j]] {
	//		leftPreorder = append(leftPreorder, preorder[j])
	//	} else {
	//		rightPreorder = append(rightPreorder, preorder[j])
	//	}
	//}

	node.Left = buildTree(leftPreorder, leftInorder)
	node.Right = buildTree(rightPreorder, rightInorder)
	return node
}