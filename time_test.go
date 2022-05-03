package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println(time.Now())
	})

	fmt.Println(time.Now())

	wg.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	// Gone Wrong

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	ticker.Stop()
	// }()

	// for time := range ticker.C {
	// 	fmt.Println(time)
	// }

	// THIS BETTER

	go func() {
		for time := range ticker.C {
			fmt.Println(time)
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Sprintln("Ticker stopped")
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	counter := 0
	for range channel {
		fmt.Println(counter)
	}
}
