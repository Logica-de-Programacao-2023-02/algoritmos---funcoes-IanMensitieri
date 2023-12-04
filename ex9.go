package main

import (
	"errors"
	"fmt"
	"strings"
)

func extrairPalavras(s string) ([]string, error) {
	if s == "" {
		return nil, errors.New("A string não pode ser vazia")
	}

	palavras := strings.Fields(s)
	return palavras, nil
}

func main() {
	texto := "Esta é uma string de exemplo"

	resultado, err := extrairPalavras(texto)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
