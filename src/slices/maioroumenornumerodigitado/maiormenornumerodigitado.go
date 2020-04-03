package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("Esperado: <número 1>,<número 2>,...,<número n> e não encontrado!\n")
		os.Exit(1)
	}
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, v := range entrada {
		valorAtual, err := strconv.Atoi(v)
		if err != nil {
			fmt.Print("%s não é um número válido!", v)
			os.Exit(1)
		}
		numeros[i] = valorAtual
	}
	maior := 0
	menor := 0

	for _, v := range numeros {
		if maior == 0 || v > maior {
			maior = v
		}
		if menor == 0 || v < menor {
			menor = v
		}
	}
	fmt.Printf("%d: Maior número \n%d: Menor número", maior, menor)

}
