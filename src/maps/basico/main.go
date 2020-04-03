package main

import "fmt"

func main() {

	capitais := map[string]string{
		"GO": "Goiania",
		"PB": "João Pessoa",
		"PR": "Curitiba"}

	fmt.Println("Tamanho: ", len(capitais))
	fmt.Println("Valores: ", capitais)
	// populando mapas

	capitais["RN"] = "Natal"
	capitais["SE"] = "Aracaju"

	fmt.Println(capitais)

	populacao := make(map[string]int)
	populacao["GO"] = 6434052
	populacao["PB"] = 3914418
	populacao["PR"] = 10997462
	populacao["RN"] = 3373960
	populacao["AM"] = 3807923
	populacao["SE"] = 2228489
	fmt.Println(populacao)

}
