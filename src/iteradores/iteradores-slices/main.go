package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1:]
	if len(entrada) <= 1 {
		fmt.Println("Esperado: <Entrada 1>, <Entrada 2>,...,<Entrada N>, mas nada encontrado!")
		os.Exit(1)
	}

	for a, valor := range entrada {
		fmt.Println("Entrada: ", a+1, " Valor: ", valor)
	}

	fmt.Println("Após alteração: ")

	for a := range entrada {
		entrada[a] = " Valor: " + strconv.Itoa(a)
		fmt.Println("Entrada: ", a+1, entrada[a])
	}

	fmt.Println("Sem acessar os valores ")

	for range entrada {
		fmt.Println("Mais uma iteração!")
	}
}
