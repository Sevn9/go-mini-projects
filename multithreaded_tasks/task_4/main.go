package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("task_4: ")

	var wg sync.WaitGroup
	var mu sync.Mutex

	var counter int

	for i := 0; i < 10; i++ {
		ctr := i
		wg.Go(func() {
			mu.Lock()
			counter++
			fmt.Println("gor: ", ctr, counter)
			mu.Unlock()
		})
	}

	wg.Wait()
}
