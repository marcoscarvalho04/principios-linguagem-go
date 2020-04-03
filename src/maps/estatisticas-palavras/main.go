package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Esperado: <Palavra 1>, <Palavra 2>,<Palavra 3>,...,<Palavra n> mas nada encontrado!")
		os.Exit(1)
	}
	palavras := os.Args[1:]
	salvarPalavras := make(map[string]int)

	for _, v := range palavras {
		valor := strings.ToUpper(string(v))
		contador, encontrado := salvarPalavras[valor]
		if encontrado {
			salvarPalavras[valor] = contador + 1
		} else {
			salvarPalavras[valor] = 1
		}
	}
	imprimir(salvarPalavras)

}

func imprimir(palavras map[string]int) {

	for palavra, contador := range palavras {
		fmt.Printf("%s = %d \n", palavra, contador)
	}
}
