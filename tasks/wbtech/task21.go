package main

import "fmt"

//	Задание 21
//	Реализовать паттерн «адаптер» на любом примере.

type ITransport interface {
	Drive()
}

type Auto struct {
	ITransport
}

func (a Auto) Drive() {
	fmt.Println("Car is driving on the road")
}

type Driver struct {
}

func (d Driver) Travel(transport ITransport) {
	transport.Drive()
}

type IAnimal interface {
	Move()
}

type Camel struct {
	IAnimal
}

func (c Camel) Move() {
	fmt.Println("Camel is move on sahara")
}

// Сам адаптер к ITransport
type CamelToTransportAdapter struct {
	ITransport
	c Camel
}

func (cta CamelToTransportAdapter) Drive() {
	cta.c.Move()
}

// Как-то так
func main() {
	d := Driver{}
	a := Auto{}

	d.Travel(a)

	c := Camel{}

	cta := CamelToTransportAdapter{c: c}

	d.Travel(cta)
}
