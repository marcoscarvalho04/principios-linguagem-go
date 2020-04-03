package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("Esperado: <número 1>, <número 2>..., <número n> e não foi encontrado!")
		os.Exit(1)
	}
	entrada := os.Args[1:]
	fmt.Println(entrada[len(entrada)/2:])
}
