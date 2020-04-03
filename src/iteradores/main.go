package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 10

	for a < b {
		a = a + 1
		fmt.Println("Iteração: ", a, " de ", b)
	} // basico.

	for i := 0; i <= 100; i++ {
		fmt.Println("Iteração: ", i, " de ", 100)
	}
}
