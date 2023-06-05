// :date 2022/4/24

// https://leetcode-cn.com/problems/maximum-subarray/

package LeetCode

func maxSubArray(nums []int) int {
	dpSum := make([]int, len(nums), len(nums))
	dpSum[0] = nums[0]
	result := dpSum[0]
	if len(nums) == 1 {
		return result
	}
	for i := 1; i < len(nums); i++ {
		tmp := dpSum[i-1] + nums[i]
		if tmp < nums[i] {
			tmp = nums[i]
		}
		dpSum[i] = tmp
	}

	for i := 1; i < len(dpSum); i++ {
		if dpSum[i] > result {
			result = dpSum[i]
		}
	}
	return result
}