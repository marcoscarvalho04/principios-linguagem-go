package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice original: ", original)

	novo := original[1:3]

	fmt.Println("Slice copiado: ", novo)

	original[2] = 13

	fmt.Println("Original pós-modificação: ", original)
	fmt.Println("Novo pós-modificação: ", novo)

	a := []string{"Paulo", "almoça", "em", "casa", "diariamente."}
	b := a[:len(a)-1]
	c := b[:len(b)-1]
	d := c[:len(c)-1]
	e := d[:len(d)-1]

	e[0] = "Thiago"

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)

}
