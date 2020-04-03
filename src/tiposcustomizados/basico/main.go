package main

import "fmt"

type ListaDeCompras []string

func main() {
	lista := make(ListaDeCompras, 6)

	lista[0] = "Alface"
	lista[1] = "Pepino"
	lista[2] = "Azeite"
	lista[3] = "Atum"
	lista[4] = "Frango"
	lista[5] = "Chocolate"
	fmt.Println(lista)

	carne, vegetais, outros := lista.Categorizar()
	fmt.Println("Carne: ", carne)
	fmt.Println("Vegetais: ", vegetais)
	fmt.Println("Outros: ", outros)
}

func (lista ListaDeCompras) Categorizar() ([]string, []string, []string) {
	var vegetais, carne, outros []string

	for _, e := range lista {
		switch e {
		case "Alface", "Pepino":
			vegetais = append(vegetais, e)
			break
		case "Atum", "Frango":
			carne = append(carne, e)
			break
		default:
			outros = append(outros, e)
			break
		}
	}
	return carne, vegetais, outros
}
