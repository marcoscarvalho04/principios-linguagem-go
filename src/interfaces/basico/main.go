package main

import (
	"fmt"
	"time"
)

type Idade struct {
	anoNascimento int
}

func (i Idade) Calcular() int {
	return time.Now().Year() - i.anoNascimento
}

type Operacao interface {
	Calcular() int
}

type Soma struct {
	operando1, operando2 int
}

type Subtracao struct {
	operando1, operando2 int
}

func (s Subtracao) Calcular() int {
	return s.operando1 - s.operando2
}

func (s Subtracao) String() string {
	return fmt.Sprintf("%d - %d", s.operando1, s.operando2)
}

func (s Soma) Calcular() int {
	return s.operando1 + s.operando2
}

func (s Soma) String() string {
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2)
}

func main() {
	operacao := make([]Operacao, 4)
	operacao[0] = Soma{10, 20}
	operacao[1] = Subtracao{30, 15}
	operacao[2] = Subtracao{10, 50}
	operacao[3] = Soma{5, 2}

	fmt.Println("Valor acumulado = ", acumulador(operacao))

	idades := make([]Operacao, 3)
	idades[0] = Idade{1969}
	idades[1] = Idade{1997}
	idades[2] = Idade{2001}

	fmt.Println("Idades acumuladas = ", acumulador(idades))
}

func acumulador(operacao []Operacao) int {

	acumulador := 0

	for _, op := range operacao {
		valor := op.Calcular()
		fmt.Printf("%v = %d\n", op, valor)
		acumulador += valor
	}
	return acumulador
}
