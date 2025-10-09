package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type responseAnswer struct {
	urlKey string
	value  string
}

func main() {
	fmt.Println("Task_8:")

	urls := []string{
		"https://www.google.com",
		"https://github.com",
		"https://www.golang.org",
		"https://error-test-urls.com",
	}

	//sendRequest(urls[0])
	resultmaps := FetchURLs(urls)

	for key, valueItem := range resultmaps {
		fmt.Println("------------------------")
		fmt.Println("\n key: " + key)
		fmt.Println("\n value: " + valueItem)
	}
}

func FetchURLs(urls []string) map[string]string {

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
			sendRequestWorker(workerId, sendRequestChannnel, resultsChannnel)
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
		results[item.urlKey] = item.value
		mu.Unlock()
	}

	return results
}

func sendRequestWorker(workerId int, urlsChan <-chan string, resultChan chan<- responseAnswer) {
	for {
		url, ok := <-urlsChan
		if !ok {
			return
		}
		if url != "" {
			fmt.Printf("worker Id=%d start working, url: %s ", workerId, url)
			resultAnswer := sendRequest(url)
			resultChan <- resultAnswer
			fmt.Printf("worker  Id=%d end working, url: %s ", workerId, url)
		} else {
			fmt.Printf("url %s is empty", url)
		}
	}
}

func sendRequest(url string) responseAnswer {

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	var answer responseAnswer

	resp, err := client.Get(url)
	if err != nil {
		answer = responseAnswer{
			urlKey: url,
			value:  "response error: " + err.Error(),
		}

		return answer
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		answer = responseAnswer{
			urlKey: url,
			value:  "cant reading responce error: " + err.Error(),
		}

		return answer
	}

	runesBody := []rune(string(body))

	if len(runesBody) > 100 {
		runesBody = runesBody[:100] // Обрезаем до 100 символов
	}

	fmt.Printf("statusCode: %d,\n body: %s", statusCode, string(runesBody))

	answer = responseAnswer{
		urlKey: url,
		value:  "statusCode: " + strconv.Itoa(statusCode) + "\n body: " + string(runesBody),
	}

	return answer
}
