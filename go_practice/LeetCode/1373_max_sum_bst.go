// :date 2022/4/6

package LeetCode


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type subTree struct {
	 MaxVal int
	 MinVal int
	 SumVal int
}

func maxSumBST(root *TreeNode) int {
	var result int
	subTreeMap := make(map[*TreeNode]*subTree)
	findSubTree(root, subTreeMap)
	for _, subTreeNode := range subTreeMap {
		if result < subTreeNode.SumVal {
			result = subTreeNode.SumVal
		}
	}
	return result
}

func findSubTree(root *TreeNode, resultMap map[*TreeNode]*subTree) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		resultMap[root] = &subTree{MaxVal: root.Val, SumVal: root.Val, MinVal: root.Val}
		return
	}

	var okLeft, okRight bool
	var leftSubNode, rightSubNode *subTree
	if root.Left != nil {
		findSubTree(root.Left, resultMap)
		leftSubNode, okLeft = resultMap[root.Left]
		if !okLeft {
			return
		}
	}
	if root.Right != nil {
		findSubTree(root.Right, resultMap)
		rightSubNode, okRight = resultMap[root.Right]
		if !okRight {
			return
		}
	}

	if (root.Left == nil || (okLeft && leftSubNode.MaxVal < root.Val)) && (root.Right == nil ||
		(okRight && root.Val < rightSubNode.MinVal)) {
		tmp := &subTree{MaxVal: root.Val, SumVal: root.Val, MinVal: root.Val}
		if okLeft {
			tmp.SumVal += leftSubNode.SumVal
			tmp.MinVal = leftSubNode.MinVal
		}
		if okRight {
			tmp.SumVal += rightSubNode.SumVal
			tmp.MaxVal = rightSubNode.MaxVal
		}
		resultMap[root] = tmp
	}
}
