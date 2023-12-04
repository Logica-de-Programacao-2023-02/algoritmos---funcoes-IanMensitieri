package main

import (
	"fmt"
)

func encontrarPosicao(slice []int, valor int) int {
	for i, elemento := range slice {
		if elemento == valor {
			return i
		}
	}
	return -1
}

func main() {
	valores := []int{10, 5, 8, 15, 7}
	valorProcurado := 8

	posicao := encontrarPosicao(valores, valorProcurado)

	fmt.Printf("A posição do primeiro elemento igual a %d é: %d\n", valorProcurado, posicao)
}
