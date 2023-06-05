// :date 2022/5/1

package LeetCode

// https://leetcode-cn.com/problems/path-sum-iii/

// prefix sum and backtrack
// https://leetcode-cn.com/problems/path-sum-iii/solution/qian-zhui-he-di-gui-hui-su-by-shi-huo-de-xia-tian/

func pathSum(root *TreeNode, target int) int {
	// current path prefix sum 
	pathPrefixSumMap := make(map[int]int)
	pathPrefixSumMap[0] = 1 // must be added that target == root.Val then first currSum - target = 0 will get 1
	return doPathSum(root, target, 0, pathPrefixSumMap)
}

func doPathSum(root *TreeNode, target, currSum int, pathPrefixSumMap map[int]int) int {
	result := 0
	if root == nil {
		return result
	}

	currSum += root.Val
	result += pathPrefixSumMap[currSum - target] // get current node path number, must do this first then add now current

	pathPrefixSumMap[currSum] = pathPrefixSumMap[currSum] + 1 // must be after pre line

	result += doPathSum(root.Left, target, currSum, pathPrefixSumMap)
	result += doPathSum(root.Right, target, currSum, pathPrefixSumMap)

	pathPrefixSumMap[currSum] = pathPrefixSumMap[currSum] - 1

	return result
}


// time usage higher，recursion search

func pathSum1(root *TreeNode, targetSum int) int {
	result := 0
	if root == nil {
		return result
	}

	result += endWithNodeNumb(root, targetSum)
	result += pathSum1(root.Left, targetSum)
	result += pathSum1(root.Right, targetSum)

	return result
}

func endWithNodeNumb(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	result := 0
	nextT := targetSum - root.Val
	if nextT == 0 {
		result += 1
	}
	result += endWithNodeNumb(root.Left, nextT)
	result += endWithNodeNumb(root.Right, nextT)

	return result
}