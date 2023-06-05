
package LeetCode

func threeSum(nums []int) [][]int {
	var result [][]int
	nMap, nIndexes := make(map[int]bool, len(nums)), make(map[int][]int)
	for i, n := range nums {
		indexes := nIndexes[n]
		indexes = append(indexes, i)
		nIndexes[n] = indexes
	}
	for i := 0; i < len(nums) - 2; i++ {
		if nMap[nums[i]] {
			continue
		}
		result = append(result, twoSum(nums[i+1:], -1*nums[i], i, nIndexes)...)
	}

	return result
}

func twoSum(nums []int, target, index int, nIndexes map[int][]int) [][]int {
	var result [][]int
	nMap := make(map[int]bool, len(nums))
	for i, n := range nums {
		nTarget := target - n
		indexes, ok := nIndexes[nTarget]
		if !ok || indexes[len(indexes)-1] <= index || indexes[len(indexes)-1] <= i+index {
			continue
		}
		if nTarget == n {
			if len(indexes) < 2 || indexes[len(indexes)-1] <= i+index {
				continue
			}
		}
		result = append(result, []int{-1*target, n, nTarget})
		nMap[n] = true
	}

	return result
}


//func threeSum(nums []int) [][]int {
//	nIndexes, PMap := make(map[int][]int, len(nums)), make(map[int]bool, len(nums)/2)
//	NMap := make(map[int]bool, len(nums)/2)
//	for i, n := range nums {
//		indexes := nIndexes[n]
//		indexes = append(indexes, i)
//		nIndexes[n] = indexes
//		if n >= 0 {
//			PMap[n] = true
//		} else {
//			NMap[n] = true
//		}
//	}
//
//	result := make([][]int, 0)
//	for target1 := range PMap {
//		fmt.Println("target1", target1)
//		if !PMap[target1] {
//			continue
//		}
//		if target1 == 0 && len(nIndexes[0]) > 2 {
//			result = append(result, []int{0, 0, 0})
//		}
//		for target2 := range NMap {
//			fmt.Println("target2", target2)
//			if !NMap[target2] {
//				continue
//			}
//			target3 := -1 * (target1 + target2)
//			indexes, ok := nIndexes[target3]
//			if !ok {
//				continue
//			}
//			fmt.Println(indexes)
//			if (target3 != target1 && target3 != target2 ) || ((target3 == target2 || target3 == target1)  && len(indexes) > 1) {
//				result = append(result, []int{target1, target2, target3})
//			}
//			if PMap[target3] {
//				PMap[target1], PMap[target3] = false, false
//			} else {
//				NMap[target2], NMap[target3] = false, false
//			}
//		}
//	}
//	return result
//}

