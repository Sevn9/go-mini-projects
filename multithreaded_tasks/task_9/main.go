package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("task_9")

	in1 := rangeGen(11, 15)
	in2 := rangeGen(21, 25)

	mergedChan := Merge(in1, in2)

	for val := range mergedChan {
		fmt.Print(val, " ")
	}
}

func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	outChan := make(chan int)

	send := func(ch <-chan int) {
		for num := range ch {
			outChan <- num
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, number := range cs {
		go send(number)
	}

	go func() {
		wg.Wait()
		close(outChan)
	}()

	return outChan
}

func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			time.Sleep(50 * time.Millisecond)
			out <- i
		}
	}()
	return out
}
