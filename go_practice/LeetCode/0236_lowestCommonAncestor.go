// :date 2022/4/30

package LeetCode

// https://leetcode-cn.com/problems/invert-binary-tree/


// Recursion contains

var ans *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	containerNode(root, p, q)

	return ans
}

func containerNode(root, p, q *TreeNode) bool {
	if root == nil {
		return  false
	}
	left := containerNode(root.Left, p, q)
	right := containerNode(root.Right, p, q)

	if (left && right) || ((root == p || root == q) && (left || right)) {
		ans = root
	}
	return left || right || root == p || root == q
}



// solution one
// find father chain
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	pf, _ := findFather(root, p, nil)
	qf, _ := findFather(root, q, nil)

	i, j := len(pf) - 1, len(qf) - 1
	var result *TreeNode
	for i >= 0 && j >= 0 {
		if pf[i] == qf[j] {
			result = pf[i]
			i--
			j--
			continue
		}
		break
	}

	return result
}

func findFather(root, node *TreeNode, result []*TreeNode) ([]*TreeNode, bool) {
	if root == nil {
		return result, false
	}
	if root == node {
		result = append(result, root)
		return result, true
	}
	var ok bool
	result, ok = findFather(root.Left, node, result)
	if ok {
		result = append(result, root)
		return result, true
	}
	result, ok = findFather(root.Right, node, result)
	if ok {
		result = append(result, root)
	}

	return result, ok
}