// :date 2021/12/7

package LeetCode

func lengthOfLongestSubstring(s string) int {
	cMap, result, start := make(map[byte]int, 26), 0, 0
	for i := 0; i < len(s); i++ {
		index, ok := cMap[s[i]]
		if ok {
			if result < i-start {
				result = i - start
			}
			newStart := index + 1
			for j := start; j < newStart; j++ {
				delete(cMap, s[j])
			}
			cMap[s[i]], start = i, newStart
			continue
		}
		cMap[s[i]] = i

		if i == len(s)-1 && i-start+1 > result {
			result = i - start + 1
		}
	}
	return result
}
