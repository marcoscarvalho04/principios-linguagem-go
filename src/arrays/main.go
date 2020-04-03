package main

import "fmt"

func main() {
	var a [3]int
	numeros := [5]int{1, 2, 3, 4, 5}
	primos := [...]int{2, 3, 5, 7, 11, 13}
	nomes := [2]string{}

	fmt.Println(a, numeros, primos, nomes)

	var multiA [2][2]int
	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, -2

	multiB := [2][2]int{{2, 13}, {-1, 16}}
	fmt.Println("Multi A", multiA)
	fmt.Println("Multi B", multiB)
}
