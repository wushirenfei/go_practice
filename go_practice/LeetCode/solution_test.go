// :date 2022/2/12

package LeetCode

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"sort"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	source1 := "abcabcbb"
	result1 := lengthOfLongestSubstring(source1)
	if result1 != 3 {
		t.Fatal("wrong out put", result1)
	}

	source2 := "pwwkew"
	result2 := lengthOfLongestSubstring(source2)
	if result2 != 3 {
		t.Fatal("wrong out put", result2)
	}

	source3 := "aab"
	result3 := lengthOfLongestSubstring(source3)
	if result3 != 2 {
		t.Fatal("wrong out put", result3)
	}

	source4 := "abcb"
	result4 := lengthOfLongestSubstring(source4)
	if result4 != 3 {
		t.Fatal("wrong out put", result4)
	}

	source5 := "tmmzuxt"
	result5 := lengthOfLongestSubstring(source5)
	if result5 != 5 {
		t.Fatal("wrong out put", result5)
	}
}

func TestLongestPalindrome(t *testing.T) {
	source1 := "babad"
	result1 := longestPalindrome(source1)
	if result1 != "bab" {
		t.Fatal("wrong out put", result1)
	}
}

func TestConvert(t *testing.T) {
	result := convert("PAYPALISHIRING", 3)
	if result != "PAHNAPLSIIGYIR" {
		t.Error("wrong result: ", result)
	}
}

func TestIsMatch(t *testing.T) {
	//fmt.Println(isMatch("aab", "c*a*b"))
	assert.Equal(t, isMatch("aab", "c*a*b"), true)
	assert.Equal(t, isMatch("aa", "a"), false)
	assert.Equal(t, isMatch("aa", "a*"), true)
	assert.Equal(t, isMatch("ab", ".*"), true)
	assert.Equal(t, isMatch("mississippi", "mis*is*p*."), false)
	assert.Equal(t, isMatch("ab", ".*c"), false)
	assert.Equal(t, isMatch("aaa", "aaaa"), false)
	assert.Equal(t, isMatch("aaa", "a*aaa"), true)
	assert.Equal(t, isMatch("aaa", "ab*a*c*a"), true)
	assert.Equal(t, isMatch("aaa", "a.a"), true)

}

func TestFindSubstring(t *testing.T) {
	result := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	fmt.Println(result)

	result2 := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	fmt.Println(result2)

	result3 := findSubstring("ababababab", []string{"ababa", "babab"})
	fmt.Println(result3)
}

func TestThreeSum(t *testing.T) {
	//for i, c := range b {
	//	fmt.Println(c, reflect.TypeOf(c))
	//	fmt.Println(i, reflect.TypeOf(b[i]))
	//}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//result := [][]int{[]int{-1, -1, 2}, []int{-1, 0 ,1}}
	//assert.Equal(t, threeSum(nums), result)

	//nums2 := []int{1, 2, -2, -1}
	//result2 := [][]int{}
	//assert.Equal(t, threeSum(nums2), result2)

	//nums3 := []int{3,0,-2,-1,1,2}
	//result3 := [][]int{[]int{-2,-1,3}, []int{-2,0,2}, []int{-1,0,1}}
	//assert.Equal(t, threeSum(nums3), result3)
}

func TestMaxArea(t *testing.T) {
	ar := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, maxArea(ar), 49)
	ar2 := []int{1, 2, 4, 3}
	assert.Equal(t, maxArea(ar2), 4)
}

func Test(t *testing.T) {
	//head := &ListNode{Val: 1}
	//head.Next = &ListNode{Val: 2}
	//fmt.Println(removeNthFromEnd(head, 2))

	a1 := []string{"()(()())", "(()())()", "()()()()", "(((())))", "((()))()", "((()()))", "()((()))", "(()(()))", "()()(())", "()(())()", "((())())", "(())()()", "(()()())"}
	a2 := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}

	sort.Strings(a1)
	sort.Strings(a2)
	fmt.Println(a1)
	fmt.Println(a2)
}

func TestNextPermutation(t *testing.T) {
	//numbs := []int{3,2,1}
	numbs := []int{1, 3, 2}
	//numbs := []int{1,2,3}
	nextPermutation(numbs)
	fmt.Println(numbs)

	//numbs2 := []int{1, 2, 7, 4, 3, 1}
	//nextPermutation(numbs2)
	//fmt.Println(numbs2)
}

func TestSlice(t *testing.T) {
	numbs2 := []int{1, 2, 7, 4, 3, 1}
	tmp2 := numbs2[6:]
	tmp1 := numbs2[1:1]
	fmt.Println(tmp2)
	fmt.Println(tmp1)
}
