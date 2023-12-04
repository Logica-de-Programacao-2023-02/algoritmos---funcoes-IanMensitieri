package main

import (
	"errors"
	"fmt"
)

func contemNumero(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice não pode estar vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numeroExemplo := 3
	sliceExemplo := []int{1, 2, 3, 4, 5}

	resultado, err := contemNumero(numeroExemplo, sliceExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
