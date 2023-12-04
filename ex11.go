package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrimo(n int) (bool, error) {
	if n < 0 {
		return false, errors.New("O número não pode ser negativo")
	}

	if n < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(n)))
	for i := 2; i <= limite; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	numeroExemplo := 13

	resultado, err := isPrimo(numeroExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
