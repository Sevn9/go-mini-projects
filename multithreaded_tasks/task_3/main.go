package main

import "fmt"

func main() {
	fmt.Println("task_3: ")
	chan1 := make(chan int)
	cancelChan := make(chan struct{})

	startGorFunc(chan1)

	for value := range chan1 {
		fmt.Println(value)
	}

	fmt.Println("select realization: ")
	//реализация с select
	chan2 := make(chan int)
	startGorFunc2(chan2, cancelChan)

	for value := range chan2 {
		fmt.Println(value)
	}
}

func startGorFunc(ch chan<- int) {
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
}

func startGorFunc2(ch chan<- int, cancel <-chan struct{}) {
	go func() {
		defer close(ch)

		//select можно использовать чтобы прервать цикл в случае отмены операции
		for i := 1; i <= 5; i++ {
			select {
			case ch <- i:
			case <-cancel:
				return
			}
		}
	}()
}
