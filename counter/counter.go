package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/* solution 1: use mutex
	var mu sync.Mutex
	count := 0
	*/
	// count := int64(0)
	var count int64

	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				/*
					mu.Lock()
					count++
					mu.Unlock()
				*/
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
