package main

import (
	"errors"
	"fmt"
)

type Pilha struct {
	valores []interface{}
}

func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("Pilha Vazia!")
	}
	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]
	return valor, nil

}

func main() {
	pilha := Pilha{}
	fmt.Println("Pilha criada com tamanho: ", pilha.Tamanho())
	fmt.Println("Vazia? ", pilha.Vazia())

	pilha.Empilhar("GO")
	pilha.Empilhar(0)
	pilha.Empilhar(5.9995)
	pilha.Empilhar("Fim")

	fmt.Println("tamanho atual da pilha: ", pilha.Tamanho())
	fmt.Println("Vazia? ", pilha.Vazia())

	for !pilha.Vazia() {
		v, err := pilha.Desempilhar()
		if err != nil {
			fmt.Println("Erro: ", err)
		} else {
			fmt.Println("Valor desempilhado: ", v)
			fmt.Println("Tamanho atual da pilha: ", pilha.Tamanho())
			fmt.Println("Vazia? ", pilha.Vazia())
		}

	}

}
