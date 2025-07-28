package main

import (
	"fmt"
	"time"
)

//	Задание 5
//	Разработать программу, которая будет последовательно отправлять значения в
//	канал, а с другой стороны канала — читать. По истечению N секунд программа
//	должна завершаться.

func reader(ch chan string, s chan bool) {
	for {
		select {
		case data, ok := <-ch:
			fmt.Printf("Данные: %s", data)

			if !ok {
				return
			}
		case <-s:
			return
		}
	}

}

func main() {
	var second int
	fmt.Println("Введите через сколько(сек) программа завершиться: ")
	_, err := fmt.Scanln(&second)
	if err != nil {
		return
	}

	d := make(chan string)
	s := make(chan bool)

	go reader(d, s)

	go func() {
		for {
			var data string
			fmt.Println("Введите данные: ")
			_, err = fmt.Scanln(&data)
			if err != nil {
				return
			}
			d <- data
		}
	}()

	time.Sleep(time.Duration(second) * time.Second)
}
