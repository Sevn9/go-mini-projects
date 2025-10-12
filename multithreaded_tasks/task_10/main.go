package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Task_10: ")

	var wg sync.WaitGroup

	//заполняем основной канал значениями
	ch := rangeGen(1, 40)

	//объединяем
	resultChanels := Split(ch, 4)

	for i, channel := range resultChanels {
		wg.Add(1)
		go func(idx int, ch <-chan int) {
			defer wg.Done()
			for item := range ch {
				fmt.Printf("Channel %d:\n", idx)
				fmt.Println(item)
			}
		}(i, channel)
	}

	wg.Wait()

}

func Split(ch <-chan int, n int) []<-chan int {
	channels := make([]chan int, n)

	// Создаем пул из n каналов
	for i := 0; i < n; i++ {
		channels[i] = make(chan int)
	}

	// Распределяет работу в круговом порядке
	// среди указанного числа каналов,
	// пока основной канал не будет закрыт.
	// При закрытии основного канала закрывает
	// все каналы и возвращается.
	toChannels := func(ch <-chan int, cs []chan int) {
		defer func(ch []chan int) {
			for _, channel := range ch {
				close(channel)
			}
		}(cs)

		//вариант 1
		for {
			for _, channel := range channels {

				val, ok := <-ch
				if !ok {
					return
				}
				channel <- val
			}
		}

		//вариант 2:
		/*
			i := 0
			for val := range ch {
				channels[i] <- val
				i = (i + 1) % n
			}
		*/

	}

	go toChannels(ch, channels)

	out := make([]<-chan int, n)

	for i, ch := range channels {
		out[i] = ch
	}
	return out
}

func rangeGen(start, stop int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			out <- i
		}
	}()

	return out
}
