package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) HelloHuman() {
	fmt.Printf("Hello, %s\n", h.Name)
}

type Action struct {
	Human
}

func (a Action) ActionHuman() {
	a.HelloHuman()
}

func main() {
	h := Human{Name: "Ass"}
	a := Action{h}

	a.ActionHuman()
}
