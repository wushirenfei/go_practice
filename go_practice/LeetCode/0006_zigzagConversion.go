// :date 2021/12/22

package LeetCode

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	rowSlice, step := make([][]byte, numRows), 2*(numRows-1)
	for i, _ := range s {
		ri := i % step
		if ri >= numRows {
			ri = step - ri
		}
		rowSlice[ri] = append(rowSlice[ri], s[i])
	}

	resultByte := make([]byte, 0)
	for _, obj := range rowSlice {
		resultByte = append(resultByte, obj...)
	}

	return string(resultByte)
}
