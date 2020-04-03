package main

import "fmt"

type Arquivo struct {
	nome       string
	tamanho    float64
	caracteres int
	palavras   int
	linhas     int
}

func (arq *Arquivo) TamanhoMedioDePalavra() float64 {
	return float64(arq.caracteres) / float64(arq.palavras)
}
func (arq *Arquivo) MediaDePalavrasPorLinha() float64 {
	return float64(arq.palavras) / float64(arq.linhas)
}

func main() {
	arquivo := Arquivo{"Teste.txt", 1234.66, 10000, 1000, 10}
	fmt.Println(arquivo)

	arquivo2 := Arquivo{nome: "Teste2.txt", tamanho: 12345}

	fmt.Println(arquivo2)

	codigoFonte := Arquivo{tamanho: 1.12, nome: "programa.go"}
	fmt.Printf("%s\t%.2fKB\n",
		codigoFonte.nome,
		codigoFonte.tamanho)
	codigoFonte.tamanho = 23.42

	fmt.Printf("%s\t%.2fKB\n",
		codigoFonte.nome,
		codigoFonte.tamanho)

	ponteiroArquivo := &Arquivo{tamanho: 7.29, nome: "arquivo.txt"}
	fmt.Printf("%s\t%.2fKB\n",
		ponteiroArquivo.nome,
		ponteiroArquivo.tamanho)

	fmt.Printf("Média de palavras por linha: %.2f\n",
		arquivo.MediaDePalavrasPorLinha())
	fmt.Printf("Tamanho médio de palavra: %.2f",
		arquivo.TamanhoMedioDePalavra())
}
