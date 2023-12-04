package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice não pode estar vazio")
	}

	sort.Strings(slice)
	resultado := strings.Join(slice, " ")

	return resultado, nil
}

func main() {
	sliceExemplo := []string{"banana", "laranja", "maçã", "abacaxi"}

	resultado, err := ordenarStrings(sliceExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
