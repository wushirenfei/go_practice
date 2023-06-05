// :date 2022/5/1

package LeetCode

// https://leetcode-cn.com/problems/diameter-of-binary-tree/

// It will be failed when result as global var because test case will be executed one by one
// that result will be set by pre case

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	findEndWithNodeMaxLen(root, &result)
	return result
}

func findEndWithNodeMaxLen(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	left := findEndWithNodeMaxLen(root.Left, result)
	right := findEndWithNodeMaxLen(root.Right, result)

	if right+left > *result {
		*result = left + right
	}

	return Max(left, right) + 1
}

func Max(a, b int) int {
	if a >= b{
		return a
	}else{
		return b
	}
}