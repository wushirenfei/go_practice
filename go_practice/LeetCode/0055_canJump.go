// :date 2022/5/6

package LeetCode

// https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	i := len(nums) - 2
	for ; i > 0; i-- {
		if nums[i] > 0 {
			continue
		}
		break
	}
	if i == 0 {
		return nums[0] > 0
	}

	j, k := i-1, 2
	for j >= 0 {
		if nums[j] < k {
			k++
			j--
			continue
		}
		break
	}
	if j < 0 {
		return false
	}
	return canJump(nums[0 : j+1])
}

// timeout
func canJump2(nums []int) bool {
	length, jumpIndex := len(nums), nums[0]
	if jumpIndex == 0 {
		return length == 1
	}

	lastIndex := length - 1
	if jumpIndex >= lastIndex {
		return true
	}
	for i := jumpIndex; i > 0; i-- {
		if canJump(nums[i:]) {
			return true
		}
	}
	return false
}
