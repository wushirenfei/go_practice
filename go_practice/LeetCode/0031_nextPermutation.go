// :date 2022/2/14

package LeetCode

func nextPermutation(nums []int)  {
	length := len(nums)
	if length == 1 {
		return
	}

	var j int
	for j = length-1; j > 0; j-- {
		if nums[j-1] < nums[j] {
			break
		}
	}
	if j == 1 && nums[0] >= nums[1] {
		j = 0
	}
	if j != 0 {
		for i := length-1; i >= j; i-- {
			if nums[j-1] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j-1]
				nums[j-1] = tmp
				break
			}
		}
	}

	start, end := j, length-1
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}

	return
}

func nextPermutation2(nums []int)  {
	length := len(nums)
	if length == 1 {
		return
	}

	var j int
	for j = length-1; j > 0; j-- {
		if nums[j-1] < nums[j] {
			break
		}
	}
	if j == 1 && nums[0] >= nums[1] {
		j = 0
	}
	if j != 0 {
		for i := length-1; i >= j; i-- {
			if nums[j-1] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j-1]
				nums[j-1] = tmp
			}
		}
	}

	start, end := j, length-1
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}

	return
}