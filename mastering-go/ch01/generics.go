package main

import (
	"fmt"
)

// with Go 1.18 Generics
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Ints := []int{1, 2, 3}
	Strings := []string{"One", "Two", "Three"}
	Print(Ints)
	Print(Strings)
}
