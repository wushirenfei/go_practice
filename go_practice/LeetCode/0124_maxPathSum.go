// :date 2022/4/25

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

package LeetCode

import "math"

var result int

func maxPathSum(root *TreeNode) int {
	result = math.MinInt32
	nodeGain(root)
	return result
}

func nodeGain(root *TreeNode) int {
	if root == nil {
		return 0
	}

	tmpLeft := nodeGain(root.Left)
	tmpRight := nodeGain(root.Right)
	if tmpLeft < 0 {
		tmpLeft = 0
	}
	if tmpRight < 0 {
		tmpRight = 0
	}

	rootSum := root.Val + tmpLeft + tmpRight
	if rootSum > result {
		result = rootSum
	}
	subGain := tmpLeft
	if subGain < tmpRight {
		subGain = tmpRight
	}
	return root.Val + subGain
}



// timeout, 93 test case passed in total 94
func maxPathSum2(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sumMap := make(map[*TreeNode]int)
	getNodeStartSum(root, sumMap)

	var leftMax, rightMax int
	corssSum, result := root.Val, root.Val
	if root.Left != nil {
		leftMax = maxPathSum2(root.Left)
		leftNodeStartSum := sumMap[root.Left]
		if leftNodeStartSum > 0 {
			corssSum += leftNodeStartSum
		}
	}
	if root.Right != nil {
		rightMax = maxPathSum2(root.Right)
		rightNodeStartSum := sumMap[root.Right]
		if rightNodeStartSum > 0 {
			corssSum += rightNodeStartSum
		}
	}
	if result < corssSum {
		result = corssSum
	}
	if root.Left != nil && result < leftMax {
		result = leftMax
	}
	if root.Right != nil && result < rightMax {
		result = rightMax
	}
	return result
}

func getNodeStartSum(root *TreeNode, sumMap map[*TreeNode]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		sumMap[root] = root.Val
		return
	}
	getNodeStartSum(root.Left, sumMap)
	getNodeStartSum(root.Right, sumMap)
	sumRoot := root.Val
	sumLeft, okLeft := sumMap[root.Left]
	sumRight, okRight := sumMap[root.Right]
	if okLeft && root.Val + sumLeft > sumRoot {
		sumRoot = root.Val + sumLeft
	}
	if okRight && root.Val + sumRight > sumRoot {
		sumRoot = root.Val + sumRight
	}
	sumMap[root] = sumRoot
}