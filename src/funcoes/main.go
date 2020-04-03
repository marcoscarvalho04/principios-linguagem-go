package main

import (
	"fmt"
)

func main() {
	imprimirDados("Fernando", 32)
	s := soma(1, 1)
	fmt.Printf("\n%d", s)
}

func imprimirDados(nome string, idade int) {
	fmt.Printf("%s possui %d anos", nome, idade)
}

func soma(a int, b int) int {
	return a + b
}
