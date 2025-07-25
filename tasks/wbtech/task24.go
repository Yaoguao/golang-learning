package main

import (
	"fmt"
	"math"
)

//	Задание 24
//	Разработать программу нахождения расстояния между двумя точками, которые
//	представлены в виде структуры Point с инкапсулированными параметрами x,y и
//	конструктором.

type Point struct {
	x, y int
}

func New(x, y int) Point {
	return Point{x, y}
}

func (p Point) Length(other Point) float64 {
	dx := float64(other.x - p.x)
	dy := float64(other.y - p.y)
	return math.Sqrt(dx*dx + dy*dy)
	//return ((other.y - p.y) * (other.y - p.y)) +
	//	((other.x - p.x) * (other.x - p.x))
}

func main() {
	p := New(0, 0)
	p2 := New(0, 4)
	fmt.Println(p.Length(p2))
}
