package main

import (
	"fmt"
)

func main() {
	s := []int{21, 22, 23}
	s = append([]int{20}, s...)

	fmt.Println(s)

	m := []int{11, 12, 13, 16, 17, 18}
	n := []int{14, 15}

	m = append(m[:3], append(n, m[3:]...)...)
	fmt.Println(m)

	a := []int{1, 2, 3, 4, 5}
	b := a[:4]

	fmt.Println("Slice 1: ", a)
	fmt.Println("Slice 2: ", b)

	c := []int{1, 2, 3, 4, 5}
	d := c[1:]

	fmt.Println("Slice 3: ", c)
	fmt.Println("Slice 4: ", d)

	e := []int{1, 2, 3, 4, 5}
	f := append(e[:2], e[3:]...)

	fmt.Println("Slice 5: ", e)
	fmt.Println("Slice 6: ", f)

	// usando função copy

	numeros := []int{1, 2, 3, 4, 5}
	dobros := make([]int, len(numeros))

	copy(dobros, numeros)

	for i := range dobros {
		dobros[i] *= 2
	}

	fmt.Println(numeros)
	fmt.Println(dobros)

}
