// :date 2022/5/3

package LeetCode

// https://leetcode-cn.com/problems/house-robber-iii/


func rob(root *TreeNode) int {
	return Max(doRob(root))
}

// It is very similar to `doRob1` but recursion too many times
func doRob(root *TreeNode) (withRoot, withoutRoot int) {
	if root == nil {
		return 0, 0
	}
	lWith, lWithout := doRob(root.Left)
	rWith, rWithout := doRob(root.Right)
	return root.Val + lWithout + rWithout, Max(lWith, lWithout) + Max(rWith, rWithout)
}



// timeout! 122 test cases are passed in 124 however it is timeout
func rob1(root *TreeNode) int {
	resInRoot := doRob1(root, true)
	resNotRoot := doRob1(root, false)
	return MaxIntSlice([]int{resInRoot, resNotRoot})
}

func doRob1(root *TreeNode, isIn bool) int {
	if root == nil {
		return 0
	}

	if isIn {
		tmpLeft := doRob1(root.Left, false)
		tmpRight := doRob1(root.Right, false)
		return root.Val + tmpLeft + tmpRight
	}
	tmpInLeft := doRob1(root.Left, true)
	tmpInRight := doRob1(root.Right, true)
	tmpNotInLeft := doRob1(root.Left, false)
	tmpNotInRight := doRob1(root.Right, false)

	withLeftWithRight := tmpInLeft + tmpInRight
	withLeftWithoutRight := tmpInLeft + tmpNotInRight
	withoutLeftWithRight := tmpNotInLeft + tmpInRight
	withoutLeftWithoutRight := tmpNotInLeft + tmpNotInRight


	return MaxIntSlice([]int{withLeftWithRight, withLeftWithoutRight, withoutLeftWithRight, withoutLeftWithoutRight})
}

func MaxIntSlice(source []int) int {
	tmp := source[0]
	for i := 1; i < len(source); i++ {
		if source[i] > tmp {
			tmp = source[i]
		}
	}

	return tmp
}


// passed, but performance is poor. It will be faster by use array to replace map
func rob2(root *TreeNode) int {
	withMap, withoutMap := make(map[*TreeNode]int), make(map[*TreeNode]int)
	doRob2(root, withMap, withoutMap)
	doRob2(root, withMap, withoutMap)
	return MaxIntSlice([]int{withMap[root], withoutMap[root]})
}

func doRob2(root *TreeNode, withMap, withoutMap map[*TreeNode]int) {
	if root == nil {
		return
	}

	doRob2(root.Left, withMap, withoutMap)
	doRob2(root.Right, withMap, withoutMap)

	withMap[root] = root.Val + withoutMap[root.Left] + withoutMap[root.Right]
	withoutMap[root] = MaxIntSlice([]int{withMap[root.Left], withoutMap[root.Left]}) +
		MaxIntSlice([]int{withMap[root.Right], withoutMap[root.Right]})
}





