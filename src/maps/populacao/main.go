package main

import (
	"fmt"
)

type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	estados := make(map[string]Estado, 6)
	estados["GO"] = Estado{"Goiás", 6434052, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

	fmt.Println(estados)

	// Recuperando valores.

	sergipe := estados["SE"]

	fmt.Println(sergipe)

	// tentar imprimir valor que não está no mapa.
	saoPaulo, encontrado := estados["SP"]
	if encontrado {
		fmt.Println(saoPaulo)
	}
	// Atualizando valores no mapa.

	estados["SE"] = Estado{"Sergipe", 2228490, "Aracaju"}

	sergipeAtualizado, encontrado := estados["SE"]

	if encontrado {
		fmt.Println(sergipeAtualizado)
	}
	// deletando de mapas

	delete(estados, "SE")

	sergipeAtualizado, encontrado = estados["SE"]

	if encontrado {
		fmt.Println(sergipeAtualizado)
	}

	// Iterando sobre mapas

	for sigla, estado := range estados {
		fmt.Printf("%s possui %d habitantes\n", sigla, estado.populacao)
	}

}
