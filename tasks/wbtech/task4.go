package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//	Реализовать постоянную запись данных в канал (главный поток). Реализовать
//	набор из N воркеров, которые читают произвольные данные из канала и
//	выводят в stdout. Необходима возможность выбора количества воркеров при
//	старте.
//	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
//	способ завершения работы всех воркеров.

func worker2(id int, dataCh <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Printf("воркер %d закрыт\n", id)
			return
		case data, ok := <-dataCh:
			if !ok {
				fmt.Printf("воркер %d: канал зак\n", id)
				return
			}
			fmt.Printf("воркер %d: обработка %d\n", id, data)
		}
	}
}

func producer(dataCh chan<- int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 0
	for {
		select {
		case <-done:
			return
		default:
			counter++
			select {
			case dataCh <- counter:
				time.Sleep(500 * time.Millisecond)
			case <-done:
				return
			}
		}
	}
}

func main() {
	var nw int
	_, err := fmt.Scan(&nw)
	if err != nil {
		return
	}

	dataCh := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	for i := 1; i <= nw; i++ {
		wg.Add(1)
		go worker2(i, dataCh, done, &wg)
	}

	wg.Add(1)
	go producer(dataCh, done, &wg)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	close(done)

	wg.Wait()
	close(dataCh)
}
