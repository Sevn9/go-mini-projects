package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	fmt.Println("task_1: ")

	//способ 1 WaitGroup
	var wg sync.WaitGroup

	//способ 1.1
	wg.Add(1)

	func() {
		defer wg.Done()
		go helloFunc()
	}()

	wg.Wait()

	//способ 1.2
	wg.Go(func() {
		go helloFunc()
	})

	wg.Wait()

	//способ 2 канал завершения
	doneChannel := make(chan struct{})
	defer close(doneChannel)

	go hello1Func(doneChannel)

	<-doneChannel

	//способ 3 atomic
	var done int32

	go func() {
		helloFunc()
		atomic.StoreInt32(&done, 1)
	}()

	for atomic.LoadInt32(&done) == 0 {
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("Main: end program")
}

func helloFunc() {
	fmt.Println("Hello from goroutine!")
}

func hello1Func(doneChannel chan<- struct{}) {
	fmt.Println("Hello from goroutine!")
	doneChannel <- struct{}{}
}
