// :date 2022/4/18

package LeetCode

// https://leetcode-cn.com/problems/unique-binary-search-trees/

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	resultMap := map[int]int{
		1: 1, 2: 2,
	}

	for tmpN := 3; tmpN <= n; tmpN++ {
		resultN := 0
		for rootN := 1; rootN <= tmpN; rootN++ {
			leftN, rightN := rootN-1, tmpN-rootN
			leftNum, rightNum := 1, 1
			if leftN > 0 {
				leftNum = resultMap[leftN]
			}
			if rightN > 0 {
				rightNum = resultMap[rightN]
			}
			resultN += leftNum * rightNum
		}
		resultMap[tmpN] = resultN
	}
	return resultMap[n]
}
