package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type responseAnswer struct {
	urlKey    string
	value     string
	errorData string
}

func main() {
	fmt.Println("Task_8:")

	urls := []string{
		"https://www.google.com",
		"https://error-test-urls.com",
		"https://github.com",
		"https://www.golang.org",
	}

	resultmaps := FetchURLs(context.Background(), urls)

	for key, valueItem := range resultmaps {
		fmt.Println("------------------------")
		fmt.Println("\n key: " + key)
		fmt.Println("\n value: " + valueItem)
	}
}

func FetchURLs(mainCtx context.Context, urls []string) map[string]string {

	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()

	var wg sync.WaitGroup

	results := make(map[string]string)

	mu := sync.Mutex{}

	workersNum := 2

	resultsChannnel := make(chan responseAnswer, workersNum)
	sendRequestChannnel := make(chan string, workersNum)

	wg.Add(workersNum)

	for i := 1; i <= workersNum; i++ {
		go func(workerId int) {
			defer wg.Done()
			sendRequestWorker(ctx, workerId, sendRequestChannnel, resultsChannnel)
		}(i)
	}

	for _, elem := range urls {
		sendRequestChannnel <- elem
	}
	close(sendRequestChannnel)

	// закрываем resultsCh после завершения воркеров
	go func() {
		wg.Wait()
		close(resultsChannnel)
	}()

	//записываем результаты с воркеров
	for item := range resultsChannnel {
		mu.Lock()
		//todo если ошибка то вызываем cancel и выходим из цикла
		if item.errorData != "" {
			results[item.urlKey] = item.errorData
			cancel()
		} else {
			results[item.urlKey] = item.value
		}
		mu.Unlock()
	}

	return results
}

func sendRequestWorker(ctx context.Context, workerId int, urlsChan <-chan string, resultChan chan<- responseAnswer) {
	for {
		select {
		case <-ctx.Done():
			return

		case url, ok := <-urlsChan:
			if !ok {
				return
			}
			if url != "" {
				fmt.Printf("worker Id=%d start working, url: %s ", workerId, url)
				resultAnswer := sendRequest(ctx, url)
				resultChan <- resultAnswer
				fmt.Printf("worker  Id=%d end working, url: %s ", workerId, url)
			} else {
				fmt.Printf("url %s is empty", url)
			}
		}
	}
}

func sendRequest(ctx context.Context, url string) responseAnswer {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	var answer responseAnswer

	//собираем запрос с контекстом
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		answer = responseAnswer{
			urlKey:    url,
			errorData: "failed to create request: " + err.Error(),
		}

		return answer
	}

	//непосредственно запрос
	resp, err := client.Do(request)

	if err != nil {
		answer = responseAnswer{
			urlKey:    url,
			errorData: "response error: " + err.Error(),
		}

		return answer
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		answer = responseAnswer{
			urlKey:    url,
			errorData: "cant reading responce error: " + err.Error(),
		}

		return answer
	}

	runesBody := []rune(string(body))

	if len(runesBody) > 100 {
		runesBody = runesBody[:100] // Обрезаем до 100 символов
	}

	fmt.Printf("statusCode: %d,\n body: %s", statusCode, string(runesBody))

	answer = responseAnswer{
		urlKey:    url,
		value:     "statusCode: " + strconv.Itoa(statusCode) + "\n body: " + string(runesBody),
		errorData: "",
	}

	return answer
}
