// :date 2021/7/12

package test

import (
	"testing"

	"go_practice/go_practice"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go_practice.Fib(30)
	}
}
