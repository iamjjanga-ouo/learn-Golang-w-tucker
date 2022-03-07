package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name string
	Age int
	Score int
	PassHitRate float32
}

type Players []Player

func (a Players) Len() int           { return len(a) }
func (a Players) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Players) Less(i, j int) bool { return a[i].Score < a[j].Score }

func main() {
	p := []Player{
		{"나통키",13,45,78.4}, 
		{"오맹태", 16,24,67.4}, 
		{"오동도", 18,54,50.8}, 
		{"황금산", 16,36,89.7},
	}

	sort.Sort(Players(p))
	fmt.Printf("%+v", p)
}