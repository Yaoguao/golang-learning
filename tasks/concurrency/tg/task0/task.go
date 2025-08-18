package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Есть список URL-адресов. Нужно написать функцию
FetchURLs(urls []string, workerCount int, timeout time.Duration) map[string]string,
которая скачивает содержимое всех страниц параллельно,
но не более workerCount запросов одновременно.

Каждый запрос должен завершаться либо
успешным ответом (строкой), либо ошибкой "timeout",если ответ не пришёл за время timeout.

Ограничения:
* Реализовать пул воркеров через каналы.
* Для таймаутов использовать context.WithTimeout или select с time.After.
* Вернуть map[URL]результат.
* Вместо HTTP-запроса можно сделать time.Sleep с разным временем, чтобы имитировать загрузку.

Пример ввода (с имитацией задержек):
urls := []string{"a.com", "b.com", "c.com"}
result := FetchURLs(urls, 2, 100*time.Millisecond)
fmt.Println(result)

Возможный вывод:
map[a.com:OK b.com:timeout c.com:OK]


go func () {

}
*/

// ОЧЕНЬ МНОГО ГОРУТИН :D

type Job struct {
	url    string
	result string
}

func Worker(timeout time.Duration, jobs chan string, results chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range jobs {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		signalCh := make(chan struct{})
		var result string

		go func(url string) {
			time.Sleep(time.Duration(rand.Int63n(150)) * time.Millisecond)

			close(signalCh)
		}(v)

		select {
		case <-ctx.Done():
			result = "timeout"
		case <-signalCh:
			result = "OK"
		}
		results <- Job{
			url:    v,
			result: result,
		}
		cancel()
	}

}

func FetchURLs(urls []string, workerCount int, timeout time.Duration) map[string]string {
	wg := &sync.WaitGroup{}
	jobs := make(chan string)
	results := make(chan Job)

	resSet := make(map[string]string, len(urls))

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go Worker(timeout, jobs, results, wg)
	}

	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for j := range results {
		resSet[j.url] = j.result
	}

	return resSet
}

func main() {
	rand.Seed(time.Now().UnixNano())
	urls := []string{"a.com", "b.com", "c.com"}
	result := FetchURLs(urls, 2, 100*time.Millisecond)
	fmt.Println(result)
}

//func Worker(ctx context.Context, chInput chan Job, url string) {
//
//	// ch struct{}
//
//	/*
//	   go {
//	       request(time sleep)
//
//	       ch struct <-
//	   }
//
//
//	   select {
//	       case ctx.Done
//
//	       case <- ch
//	   }
//	*/
//}
