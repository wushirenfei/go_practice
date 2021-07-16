// :date 2021/7/12

package test

import (
	"testing"

	"go_practice/go_practice"
)

// go test -v atomic_test.go -test.bench ReadWriteWith -test.run ReadWriteWith

func BenchmarkAtomic(test *testing.B) {
	go_practice.AtomicPractice()
}

func TestAtomic(test *testing.T) {
	go_practice.AtomicPractice()
}

func BenchmarkRandInt(t *testing.B) {
	go_practice.GenIntByRandInt()
}

func BenchmarkNewRandInt(t *testing.B) {
	go_practice.GenIntByNewRandInt()
}

func BenchmarkReadWriteWithMutex(t *testing.B) {
	//go_practice.ReadWriteWithMutex(8, 2)
	go_practice.ReadWriteWithMutex(5, 5)
}

func BenchmarkReadWriteWithAtomic(t *testing.B) {
	//go_practice.ReadWriteWithAtomic(8, 2)
	go_practice.ReadWriteWithAtomic(5, 5)
}
