package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano()) //для составления рандомных чисел. больше тебе знать не нужно
}

// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000) //рандомное число
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обертку predictableFunc, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
func predictableFunc() (int64, error) {
	resCh := make(chan int64)
	// timer := time.NewTimer(time.Second) // можно использовать таймер
	go func() {
		resCh <- unpredictableFunc()
	}()

	select {
	case v := <-resCh:
		return v, nil
	case <-time.After(time.Second):
		return 0, errors.New("ERROR")
	case <-time.After(time.Second):
		return 0, errors.New("ERROR")
	}
}

//	С использованием контекста
//func predictableFunc() (int64, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	resCh := make(chan int64)
//
//	go func() {
//		resCh <- unpredictableFunc()
//	}()
//
//	select {
//	case v := <-resCh:
//		return v, nil
//	case <-ctx.Done():
//		return 0, errors.New("ERROR")
//	}
//}

func main() {
	fmt.Println("started")

	fmt.Println(predictableFunc())
}
