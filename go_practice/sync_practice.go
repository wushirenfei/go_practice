// :date 2021/7/13

package go_practice

import (
	"sync"
	"time"
)

func ReadWriteMap() {
	pool := make(map[string]int)
	var wg sync.WaitGroup
	var locker sync.RWMutex
	keys := []string{"key1", "key2", "key3", "key4"}
	for i, key := range keys {
		wg.Add(1)
		go func() {
			defer wg.Done()
			locker.Lock()
			defer locker.Unlock()
			pool[key] = i
			time.Sleep(1 * time.Microsecond)
		}()
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//locker.RLock()
			//defer locker.RUnlock()
			locker.Lock()
			defer locker.Unlock()
			_ = pool[keys[i%4]]
			time.Sleep(1 * time.Microsecond)
		}()
	}
	wg.Wait()
}

func ReadWriteSyncMap() {
	var pool sync.Map
	var wg sync.WaitGroup
	keys := []string{"key1", "key2", "key3", "key4"}
	for i, key := range keys {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pool.Store(key, i)
			time.Sleep(1 * time.Microsecond)
		}()
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = pool.Load(keys[i%4])
			time.Sleep(1 * time.Microsecond)
		}()
	}
	wg.Wait()
}
