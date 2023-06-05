// :date 2021/12/8

package LeetCode

func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 || length == 1 {
		return s
	}

	window := length
	for ; window > 1; window-- {
		for i := 0; i+window <= length; i++ {
			flag := true
			for start, end := i, i+window-1; start < end; {
				if s[start] != s[end] {
					flag = false
					break
				}
				start++
				end--
			}
			if flag {
				return s[i : i+window]
			}
		}
	}

	return string(s[0])
}

//func longestPalindrome(s string) string {
//	cMap, start, resultStart, resultEnd := make(map[byte]int), 0, 0, 0
//	for i := 0; i < len(s); i++ {
//		index, ok := cMap[s[i]]
//		if ok {
//			if i-index >= resultEnd-resultStart {
//				resultStart, resultEnd = index, i
//			}
//			for j := start; j < index+1; j++ {
//				delete(cMap, s[j])
//			}
//			start = index + 1
//			continue
//		}
//		cMap[s[i]] = i
//	}
//	return s[resultStart : resultEnd+1]
//}
