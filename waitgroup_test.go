package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func RunAsynchronous(group *sync.WaitGroup, counter *int, mutex *sync.Mutex) {
	defer group.Done()

	group.Add(1)

	mutex.Lock()

	fmt.Println("Hai ke", int(*counter))
	*counter = *counter + 1

	mutex.Unlock()
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	var mutex sync.Mutex
	counter := 0

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group, &counter, &mutex)
	}

	group.Wait()
	fmt.Println("Complete")
}
