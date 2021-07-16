// :date 2021/7/13

package test

import (
	"testing"

	"go_practice/go_practice"
)

func BenchmarkReadWriteMap(t *testing.B) {
	go_practice.ReadWriteMap()
}

func BenchmarkReadWriteSyncMap(t *testing.B) {
	go_practice.ReadWriteSyncMap()
}
