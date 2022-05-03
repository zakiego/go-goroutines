package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	var counter int64 = 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				fmt.Println(counter)
				atomic.AddInt64(&counter, 1)
				// counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}
