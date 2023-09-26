package main

import "fmt"

func main() {
	numGenerator := generator()

	for i := 0; i < 5; i++ {
		fmt.Println(numGenerator(), "\t")
	}
}

// generator returns a function that returns an interger
func generator() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
