package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	/* Solution 1: mutex
	var mu sync.Mutex
	count := 0
	*/

	count := int64(0)
	nGR, nIter := 10, 1_000

	var wg sync.WaitGroup

	wg.Add(nGR)
	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				// Solution 2: sync/atomic
				atomic.AddInt64(&count, 1)
				/* Solution 1: mutex
				mu.Lock()
				count++
				mu.Unlock()
				*/
				/*
					count ++ is translated to:
					fetch count
					increment count
					store count
				*/
				time.Sleep(time.Microsecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}

/*
go run -race ./count

"-race" is supportd by
- run
- build
- test

Rule of thumb: use "go test -race"
*/
