package main

import "fmt"

// 1. 8

// 2
type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, HP int, Speed float64) *Actor {
	return &Actor{name, HP, Speed}
}

func main() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}

// 3 : 1개
