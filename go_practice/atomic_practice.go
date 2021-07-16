// :date 2021/7/12

package go_practice

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var locker sync.RWMutex

func AtomicPractice() {
	var a = 1
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 1000; i++ {
			locker.Lock()
			a++
			locker.Unlock()
		}
	}()

	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 50; i++ {
			locker.RLock()
			fmt.Println("read ", i, "a as", a)
			locker.RUnlock()
		}

	}()
	wg.Wait()
}

func GenIntByRandInt() {
	numb := 100
	var wg sync.WaitGroup
	wg.Add(numb)
	for i := 0; i < numb; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			for j := 0; j < 10000; j++ {
				rand.Int()
			}
		}()
	}
	wg.Wait()
}

func GenIntByNewRandInt() {
	numb := 100
	var wg sync.WaitGroup
	wg.Add(numb)
	for i := 0; i < numb; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			newRand := rand.New(rand.NewSource(time.Now().Unix()))
			for j := 0; j < 10000; j++ {
				newRand.Int()
			}
		}()
	}
	wg.Wait()
}

func ReadWriteWithMutex(readNumb, writeNumb int) {
	var locker sync.RWMutex
	var wg sync.WaitGroup
	count := 0

	for i := 0; i < writeNumb*1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			locker.Lock()
			defer locker.Unlock()
			count++
			time.Sleep(1 * time.Microsecond)
		}()
	}

	for j := 0; j < readNumb*1000; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			locker.RLock()
			defer locker.RUnlock()
			_ = count
			time.Sleep(1 * time.Microsecond)
		}()
	}
	wg.Wait()
}

func ReadWriteWithAtomic(readNumb, writeNumb int) {
	var wg sync.WaitGroup
	var count int32

	for i := 0; i < writeNumb*1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
			atomic.AddInt32(&count, 1)
			time.Sleep(1 * time.Microsecond)
		}()
	}

	for j := 0; j < readNumb*1000; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			locker.RLock()
			defer locker.RUnlock()
			_ = atomic.LoadInt32(&count)
			time.Sleep(1 * time.Microsecond)
		}()
	}
	wg.Wait()
}
