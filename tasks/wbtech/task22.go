package main

import (
	"fmt"
	"math/big"
)

//	Задание 22
//	Разработать программу, которая перемножает, делит, складывает, вычитает две
//	числовых переменных a,b, значение которых > 2^20.

func calculate(a, b *big.Int, op string) *big.Int {
	switch op {
	case "+":
		return new(big.Int).Add(a, b)
	case "-":
		return new(big.Int).Sub(a, b)
	case "*":
		return new(big.Int).Mul(a, b)
	case "/":
		return new(big.Int).Div(a, b)
	default:
		return nil
	}
}

func main() {
	a := big.NewInt(1 << 60) // 2^30
	b := big.NewInt(1 << 50) // 2^20
	result := calculate(a, b, "*")
	fmt.Println(result)
}
