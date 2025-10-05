package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// инициализация канала
	input := make(chan int)

	//через 10 сек прекратить выполнение программы
	workTime := 10 * time.Second

	/* создание контекста  */
	ctx, cancel := context.WithTimeout(context.Background(), workTime)
	defer cancel()

	wg.Go(func() {
		StartBatchProcessor(ctx, input)
	})

	// сбор данных
	wg.Go(func() {
		for i := 1; i <= 20; i++ {
			input <- i
			time.Sleep(300 * time.Millisecond)

			//через cancel вручную
			/*
				cancel()
				close(input)
				break
			*/
		}
	})

	wg.Wait()
	fmt.Println("Main: processing stopped")

	var counter int32 = 0

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			atomic.AddInt32(&counter, 1)
		})
	}

	wg.Wait()
}

func StartBatchProcessor(ctx context.Context, input <-chan int) {

	for {
		//ставим таймер
		workTime := 2 * time.Second
		timer := time.NewTimer(workTime)

		var batch []int

	outer:
		for {
			select {
			// если контекст закрылся по главному таймеру или сработала отмена в main
			case <-ctx.Done():
				batchWorker(batch)
				timer.Stop()
				return

			//внутренний таймер в 2 сек сработал
			case <-timer.C:
				//обрабатываем что успели собрать
				batchWorker(batch)
				//выходим из for
				break outer

			case value, ok := <-input:
				if !ok {
					batchWorker(batch)
					timer.Stop()
					return
				}
				batch = append(batch, value)

				if len(batch) == 5 {
					batchWorker(batch)
					//обнуляем слайс
					batch = batch[:0]
				}
			}
		}
	}
}

func batchWorker(batch []int) {
	//обработка батча
	fmt.Println("Processed batch:", batch)
}
