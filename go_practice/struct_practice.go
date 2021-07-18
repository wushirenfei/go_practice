// :date 2021/7/18

package go_practice

import (
	"fmt"
)

type IntG int

func (a IntG) Less(b IntG) bool {
	return a < b
}

func (a *IntG) Add(b IntG) IntG {
	*a += b
	return *a
}

// 定义接口
type LessAdder interface {
	Less(b IntG) bool
	Add(b IntG)
}

type A struct {
	attr int
}

func (a A) Call() {
	fmt.Println("call a")
	a.attr = 100
}

type AInterface interface {
	Call()
}

func RunStructDemo() {
	a := A{}
	fmt.Println(a)
	i := a
	i.Call()
	fmt.Println(a)
}
