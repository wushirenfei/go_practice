// :date 2022/2/10

package LeetCode

func maxAreaTimeout(height []int) int {
	result := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			high := height[i]
			if high > height[j] {
				high = height[j]
			}
			area := high * (j - i)
			if area > result {
				result = area
			}
		}
	}
	return result
}

func maxArea(height []int) int {
	result, head, tail := 0, 0, len(height)-1
	for head < tail {
		high, delta := height[head], tail-head
		if height[head] >= height[tail] {
			high = height[tail]
			newTail := tail - 1
			for {
				if newTail > head && height[tail] >= height[newTail] {
					newTail--
					continue
				}
				break
			}
			tail = newTail
		} else {
			newHead := head + 1
			for {
				if tail > newHead && height[head] >= height[newHead] {
					newHead++
					continue
				}
				break
			}
			head = newHead
		}
		area := high * delta
		if area > result {
			result = area
		}
	}
	return result
}
