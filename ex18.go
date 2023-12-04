package main

import (
	"errors"
	"fmt"
)

func aplicarFuncaoESomar(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice não pode estar vazio")
	}

	soma := 0
	for _, elemento := range slice {
		resultado := funcao(elemento)
		soma += resultado
	}

	return soma, nil
}

func main() {
	sliceExemplo := []int{1, 2, 3, 4, 5}

	funcaoExemplo := func(x int) int {
		return x * x
	}

	resultado, err := aplicarFuncaoESomar(sliceExemplo, funcaoExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
