// :date 2022/5/6

package LeetCode

import (
	"math"
)

// https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	Stack    []int
	TopIndex int
	MinValue int
}

func Constructor() MinStack {
	return MinStack{
		Stack:    make([]int, 0, 16),
		TopIndex: -1,
		MinValue: math.MaxInt,
	}
}

func (s *MinStack) Push(val int) {
	s.Stack = append(s.Stack, val)
	s.TopIndex += 1
	if val < s.MinValue {
		s.MinValue = val
	}

	a := make(map[int]int)
	delete(a, 1)
}

func (s *MinStack) Pop() {
	if s.TopIndex == 0 {
		s.Stack, s.TopIndex, s.MinValue = make([]int, 0, 16), -1, math.MaxInt
		return
	}
	if s.TopIndex > 0 {
		if s.Stack[s.TopIndex] == s.MinValue {
			s.MinValue = s.Stack[0]
			for i := 1; i < s.TopIndex; i++ {
				if s.Stack[i] < s.MinValue {
					s.MinValue = s.Stack[i]
				}
			}
		}
		s.Stack = s.Stack[:s.TopIndex]
		s.TopIndex -= 1
	}
}

func (s *MinStack) Top() int {
	return s.Stack[s.TopIndex]
}

func (s *MinStack) GetMin() int {
	return s.MinValue
}
