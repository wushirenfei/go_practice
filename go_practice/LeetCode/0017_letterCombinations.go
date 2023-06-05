// :date 2022/2/12

package LeetCode

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	numbMap := make(map[uint8][]string, 8)
	numbMap['2'] = []string{"a", "b", "c"}
	numbMap['3'] = []string{"d", "e", "f"}
	numbMap['4'] = []string{"g", "h", "i"}
	numbMap['5'] = []string{"j", "k", "l"}
	numbMap['6'] = []string{"m", "n", "o"}
	numbMap['7'] = []string{"p", "q", "r", "s"}
	numbMap['8'] = []string{"t", "u", "v"}
	numbMap['9'] = []string{"w", "x", "y", "z"}

	for i, _ := range digits {
		tmp := numbMap[digits[i]]
		if len(result) == 0 {
			result = tmp
			continue
		}
		tmp2 := make([]string, 0, len(result) * len(tmp))
		for _, tc := range tmp {
			for _, tr := range result {
				tmp2 = append(tmp2, tr+tc)
			}
		}
		result = tmp2
	}
	return result
}