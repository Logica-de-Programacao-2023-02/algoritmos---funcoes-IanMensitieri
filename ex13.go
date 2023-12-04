package main

import (
	"errors"
	"fmt"
)

func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("O número não pode ser negativo")
	}

	soma := 0

	for numero > 0 {
		digito := numero % 10
		soma += digito
		numero /= 10
	}

	return soma, nil
}

func main() {
	numeroExemplo := 12345

	resultado, err := somaDigitos(numeroExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
