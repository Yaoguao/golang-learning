package main

import (
	"fmt"
	"time"
)

/*
Реализуйте три функции printFirst(), printSecond(), printThird() таким образом, чтобы при вызове из разных горутин они всегда выполнялись строго в порядке: First, затем Second, затем Third. Используйте средства синхронизации (не time.Sleep).

Давай условимся, что все три функции будут обязательны вызыван, просто их порядок вызова может быть разным

Пример:
Вывод всегда:
first
second
third



*/

type OrderedPrinter struct {
	secondSignal, thirdSignal chan struct{}
}

func NewOrderedPrinter() *OrderedPrinter {
	return &OrderedPrinter{
		secondSignal: make(chan struct{}),
		thirdSignal:  make(chan struct{}),
	}
}

func (op *OrderedPrinter) printFirst(print func()) {
	print()
	op.secondSignal <- struct{}{}
}

func (op *OrderedPrinter) printSecond(print func()) {
	select {
	case <-op.secondSignal:
		print()
		op.thirdSignal <- struct{}{}
		return
	}
}

func (op *OrderedPrinter) printThird(print func()) {
	select {
	case <-op.thirdSignal:
		print()
		return
	}
}

/*
func (op *OrderedPrinter) printFirst(print func()) {
	print()
	close(op.secondSignal)
}

func (op *OrderedPrinter) printSecond(print func()) {
	<-op.secondSignal
	print()
	close(op.thirdSignal)
}

func (op *OrderedPrinter) printThird(print func()) {
	<-op.thirdSignal
	print()
}

*/

func main() {
	op := NewOrderedPrinter()

	go op.printSecond(func() { fmt.Println("second") })
	go op.printThird(func() { fmt.Println("third") })
	go op.printFirst(func() { fmt.Println("first") })

	time.Sleep(1 * time.Second)

}
