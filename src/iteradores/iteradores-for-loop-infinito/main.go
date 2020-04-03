package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 0
	for {
		n++
		i := rand.Intn(42)
		fmt.Println("Iterando: ", n, " vezes. Valor aleatório gerado: ", i)

		if i%42 == 0 {
			fmt.Println(i, " é divisível por 42, saindo")
			break
		}
	}
	fmt.Println("Saindo após: ", n, " iterações")
}
