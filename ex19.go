package main

import (
	"errors"
	"fmt"
)

func obterPrimosAteN(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("O número não pode ser negativo")
	}

	var primos []int

	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func isPrimo(n int) bool {
	if n < 2 {
		return false
	}

	limite := int(n/2) + 1
	for i := 2; i < limite; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	numeroExemplo := 20

	resultado, err := obterPrimosAteN(numeroExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
