package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	texto := "Anderson tem 21 anos"
	expr := regexp.MustCompile("\\d")

	fmt.Println(expr.ReplaceAllString(texto, "3"))
	transformadora :=
		func(s string) string {
			return strings.ToUpper(s)
		}

	expr = regexp.MustCompile("\\b\\w")
	texto = "antonio carlos jobim"

	processado := expr.ReplaceAllStringFunc(texto, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(processado)
	processadoVariavelFuncao := expr.ReplaceAllStringFunc(texto, transformadora)
	fmt.Println(processadoVariavelFuncao)
}
