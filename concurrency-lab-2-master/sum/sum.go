package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}

	sum := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			sum = sum + 1
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
