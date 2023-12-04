package main

import (
	"errors"
	"fmt"
)

func stringsComMaisDe5Caracteres(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice não pode estar vazio")
	}

	var resultado []string
	for _, s := range slice {
		if len(s) > 5 {
			resultado = append(resultado, s)
		}
	}

	return resultado, nil
}

func main() {
	sliceExemplo := []string{"apple", "banana", "orange", "kiwi", "pear"}

	resultado, err := stringsComMaisDe5Caracteres(sliceExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
