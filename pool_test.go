package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var wg sync.WaitGroup

	pool := sync.Pool{
		New: func() interface{} {
			return "Blank"
		},
	}

	pool.Put("Muhammad")
	pool.Put("Zakiyuddin")
	pool.Put("Munziri")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			data := pool.Get()
			fmt.Println(data)

			time.Sleep(1 * time.Second)

			pool.Put(data)
		}()
	}

	wg.Wait()
	fmt.Println("Selesai")
}
