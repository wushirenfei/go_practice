// :date 2021/12/24
// 给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/regular-expression-matching

package LeetCode

//func isMatch(s string, p string) bool {
//	var star, point uint8
//	star, point = "*"[0], "."[0]
//	if len(p) == 0 {
//		return len(s) == 0
//	}
//	if len(p) > 1 && p[1] == star {
//		return isMatch(s, p[2:]) || (len(s) > 0 && (s[0] == p[0] || p[0] == point) && isMatch(s[1:], p))
//	}
//	return len(s) > 0 && (s[0] == p[0] || p[0] == point) && isMatch(s[1:], p[1:])
//}

//We define dp[i][j] to be true if s[0..i) matches p[0..j) and false otherwise. The state equations will be:
//
//dp[i][j] = dp[i - 1][j - 1], if p[j - 1] != '*' && (s[i - 1] == p[j - 1] || p[j - 1] == '.');
//dp[i][j] = dp[i][j - 2], if p[j - 1] == '*' and the pattern repeats for 0 time;
//dp[i][j] = dp[i - 1][j] && (s[i - 1] == p[j - 2] || p[j - 2] == '.'), if p[j - 1] == '*' and the pattern repeats for at least 1 time.

func isMatch(s string, p string) bool {
	sl, pl := len(s), len(p)
	dp := make([][]bool, sl+1)
	for i := range dp {
		dp[i] = make([]bool, pl+1)
	}

	dp[0][0] = true
	for i := 0; i <= sl; i++ {
		for j := 1; j <= pl; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				dp[i][j] = i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') && dp[i-1][j-1]
			}
		}
	}
	return dp[sl][pl]
}

//func isMatch(s string, p string) bool {
//	var star, point uint8
//	star, point = "*"[0], "."[0]
//	switch len(p) {
//	case 0:
//		return len(s) == 0
//	case 1:
//		return len(s) == 1 && (s[0] == p[0] || p[0] == point)
//	default:
//
//		if p[1] != star {
//			if len(s) > 0 && (s[0] == p[0] || p[0] == point) {
//				return isMatch(s[1:], p[1:])
//			}
//			return false
//		} else {
//			if len(s) > 0 && (s[0] == p[0] || p[0] == point) {
//				if isMatch(s, p[2:]) == true {
//					return true
//				} else {
//					return isMatch(s[1:], p)
//				}
//			} else {
//				return isMatch(s, p[2:])
//			}
//		}
//	}
//}
