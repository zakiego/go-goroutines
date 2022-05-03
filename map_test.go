package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, wg *sync.WaitGroup) {
	defer wg.Done()

	data.Store(value, value)
}

func TestMap(t *testing.T) {
	wg := &sync.WaitGroup{}
	data := &sync.Map{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go AddToMap(data, i, wg)
	}

	wg.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
