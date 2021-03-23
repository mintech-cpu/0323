package main

import (
	"fmt"
)

// 構造体とスライス
// Goで頻出

type Person struct {
	Name string
}

func main() {
	ps := make([]Person, 5)
	fmt.Println(ps)

	ps[0].Name = "Mike"
	ps[1].Name = "Ken"
	ps[2].Name = "Ron"
	ps[3].Name = "Sam"
	ps[4].Name = "Jon"
	fmt.Println(ps)

}
