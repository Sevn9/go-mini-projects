package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("task_2: ")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(
			func() {
				printId(i)
			})
	}

	wg.Wait()
	fmt.Println("done with sync.WaitGroup")

	//решение без WaitGroup
	doneChannel := make(chan struct{}, 5)

	for i := 1; i <= 5; i++ {
		go func(chan1 chan<- struct{}) {
			printId2(i, chan1)
		}(doneChannel)
	}

	for i := 1; i <= 5; i++ {
		<-doneChannel
	}

	fmt.Println("done with channel")

}

func printId(id int) {
	fmt.Println("id: ", id)
}

func printId2(id int, channel chan<- struct{}) {
	fmt.Println("id: ", id)
	channel <- struct{}{}
}
