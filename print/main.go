package main

import (
	"fmt"
)

func main() {
	string1 := "This is a string"

	for _, v := range string1 {
		fmt.Printf("Unicode code point : %U - character '%c' - binary %b - hex %X\n", v, v, v, v)
	}
}
