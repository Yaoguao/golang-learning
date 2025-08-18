package main

import "sync"

// Напишите функцию, которая мержит n-каналов в один
/*


 [] [] [] [] [] - chan

go func {
	select {
	case <- []
	case <- []
	}

}

*/

func merge(channels ...chan int64) chan int64 {
	wg := sync.WaitGroup{}
	totalCh := make(chan int64)

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch chan int64) {
			defer wg.Done()
			for v := range ch {
				totalCh <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(totalCh)
	}()

	return totalCh
}

func main() {
	channels := make([]chan int64, 10)
	for i := range channels {
		channels[i] = make(chan int64)
	}

	for i := range channels {
		go func(i int) {
			channels[i] <- int64(i)
			close(channels[i])
		}(i)
	}

	for v := range merge(channels...) {
		println(v)
	}
}
