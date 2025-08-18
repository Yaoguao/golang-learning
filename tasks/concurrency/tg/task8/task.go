package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Вы разрабатываете сервис, который обрабатывает изображения.
// Каждое изображение проходит дорогостоящую обработку, например, наложение подяного знака.
// Поскольку обработка кажлдого изображения занимает значительное время,
// необходимо обрабатывать их параллельно, чтобы ускорить процесс.
// Однако, чтобы избежать излишней нагрузки на систему, вы хотите ограничить количество
// одновременно работающих горутин.

type Job struct {
	image  string
	result string
}

func Worker(timeout time.Duration, images chan string, results chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for image := range images {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		signalCh := make(chan struct{})
		var result string

		go func() {
			time.Sleep(time.Duration(rand.Int63n(450)) * time.Millisecond)

			close(signalCh)
		}()

		select {
		case <-ctx.Done():
			result = "timeout"
		case <-signalCh:
			result = "OK"
		}
		results <- Job{
			image:  image,
			result: result,
		}
		cancel()
	}

}

func ProcessingImages(images []string, workerCount int, timeout time.Duration) map[string]string {
	wg := &sync.WaitGroup{}
	jobs := make(chan string)
	results := make(chan Job)

	resSet := make(map[string]string, len(images))

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go Worker(timeout, jobs, results, wg)
	}

	go func() {
		for _, image := range images {
			jobs <- image
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for j := range results {
		resSet[j.image] = j.result
	}

	return resSet
}

func main() {
	rand.Seed(time.Now().UnixNano())
	images := []string{
		"image1.png",
		"image2.png",
		"image3.png",
	}
	result := ProcessingImages(images, 2, 200*time.Millisecond)
	fmt.Println(result)
}
