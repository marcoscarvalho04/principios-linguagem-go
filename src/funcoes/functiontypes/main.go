package main

import "fmt"

type Agregadora func(n, m int) int

func Agregar(
	valores []int,
	inicial int,
	fn Agregadora,
) int {
	agregado := inicial

	for _, v := range valores {
		agregado = fn(v, agregado)
	}
	return agregado
}

func CalcularSoma(valores []int) int {
	soma := func(n, m int) int {
		return n + m
	}
	return Agregar(valores, 0, soma)
}

func CalcularProdutos(valores []int) int {
	produto := func(n, m int) int {
		return n * m
	}
	return Agregar(valores, 1, produto)
}

func main() {
	valores := []int{-3, -2, 5, 7, 8, 22, 32, -1}

	fmt.Println(CalcularProdutos(valores))
	fmt.Println(CalcularSoma(valores))

}
