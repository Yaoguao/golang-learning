package main

import (
	"fmt"
)

// Вы разрабатываете систему для обработки финансовых транзакций.
// Каждая транзакция проходит несколько этапов обработки:
// 1. Чтение транзакций из исходных данных.
// 2. Фильтрация транзакций: убирает транзакции с отрицательными суммами.
// 3. Конвертация валюты: преобразуем сумму транзакции в доллары.
// 4. Сохранение результатов: записываем обработанные транзакции в итоговый список.

/*

readTransaction(... sum) <- chan int

filterTransaction(<- chan int) <- chan int

convertTransaction(<- chan int) <- chan int

saveTransaction(<- chan int)

*/

func ReadTransaction(sum ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, val := range sum {
			ch <- val
		}
		close(ch)
	}()

	return ch
}

func FilterTransaction(tranInput <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for n := range tranInput {
			if n > 0 {
				ch <- n
			}
		}
		close(ch)
	}()

	return ch
}

func ConvertTransaction(tranInput <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for n := range tranInput {
			ch <- n / 90
		}
		close(ch)
	}()

	return ch
}

func SaveTransaction(tranInput <-chan int) {
	lst := make([]int, len(tranInput))

	for n := range tranInput {
		lst = append(lst, n)
	}

	fmt.Println(lst)
}

func main() {
	SaveTransaction(
		ConvertTransaction(
			FilterTransaction(
				ReadTransaction(500, 400, 600, 360, -23, 56))))
}
