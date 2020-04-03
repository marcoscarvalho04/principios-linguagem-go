package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	palavras := os.Args[1:]
	estatiticas := colherEstatisticas(palavras)
	imprimir(estatiticas)
}

func colherEstatisticas(palavras []string) map[string]int {
	estatiticas := make(map[string]int)
	for _, v := range palavras {
		inicial := strings.ToUpper(string(v[0]))
		contador, encontrado := estatiticas[inicial]
		if encontrado {
			estatiticas[inicial] = contador + 1
		} else {
			estatiticas[inicial] = 1
		}
	}
	return estatiticas
}

func imprimir(estatiticas map[string]int) {
	fmt.Printf("Contagem de palavras iniciadas em: \n")
	for inicial, contador := range estatiticas {
		fmt.Printf("%s = %d \n", inicial, contador)
	}
}
