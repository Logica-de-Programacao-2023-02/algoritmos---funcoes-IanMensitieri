package main

import (
	"errors"
	"fmt"
	"strings"
)

func substituirVogaisPorAsterisco(s string) (string, error) {
	if s == "" {
		return "", errors.New("A string não pode ser vazia")
	}

	resultado := strings.Map(func(r rune) rune {
		if strings.ContainsRune("aeiouAEIOU", r) {
			return '*'
		}
		return r
	}, s)

	return resultado, nil
}

func main() {
	stringExemplo := "Hello World"

	resultado, err := substituirVogaisPorAsterisco(stringExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
