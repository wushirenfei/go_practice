// :date 2021/12/23

package LeetCode

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	result := make([]int, 0)
	lenS, lenW, lenWS := len(s), len(words[0]), len(words)
	if lenS < lenW*lenWS {
		return result
	}

	wordRemainMap, count := make(map[string]int, len(words)), 0
	initCounter := func() {
		for _, word := range words {
			wordRemainMap[word] = 0
		}
		for _, word := range words {
			wordRemainMap[word]++
		}
		count = 0
	}

	for index := 0; index < lenW; index++ {
		j := index
		initCounter()
		moveNext := func() {
			wordRemainMap[s[j:j+lenW]]++
			count--
			j += lenW

		}
		for j < lenS-lenWS*lenW+1 {
			word := s[j+count*lenW : j+(count+1)*lenW]
			remain, ok := wordRemainMap[word]
			switch {
			case !ok:
				j = j + (count+1)*lenW
				if count != 0 {
					initCounter()
				}
			case remain == 0:
				moveNext()
			default:
				wordRemainMap[word]--
				count++
				if count == lenWS {
					result = append(result, j)
					moveNext()

				}
			}
		}
	}
	return result
}

//func findSubstring(s string, words []string) []int {
//	result := make([]int, 0)
//	wordLen, wordsLen := len(words[0]), len(words)
//	if wordsLen == 0 || len(s) < wordLen*wordsLen {
//		return result
//	}
//
//	wordsMap := make(map[string]int, len(words))
//	initWordsMap := func() {
//		for _, word := range words {
//			wordsMap[word] = 0
//		}
//		for _, word := range words {
//			wordsMap[word] += 1
//		}
//	}
//
//	for i := 0; i < len(s)-wordLen*wordsLen+1; i += wordLen {
//		initWordsMap()
//		flagIndex := true
//		for j := i; j < i+wordLen*wordsLen; j += wordLen {
//			tmpStr := string(s[j : j+wordLen])
//			wordsMap[tmpStr] -= 1
//			if wordsMap[tmpStr] < 0 {
//				flagIndex = false
//				break
//			}
//		}
//		for _, v := range wordsMap {
//			if v > 0 {
//				flagIndex = false
//				break
//			}
//		}
//		if flagIndex {
//			result = append(result, i)
//		}
//	}
//
//	return result
//}
