package main

import (
	"errors"
	"fmt"
)

func intersecaoSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Ambos os slices devem estar preenchidos")
	}

	mapeamento := make(map[int]bool)
	for _, num := range slice1 {
		mapeamento[num] = true
	}

	var resultado []int
	for _, num := range slice2 {
		if mapeamento[num] {
			resultado = append(resultado, num)
		}
	}

	return resultado, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	resultado, err := intersecaoSlices(slice1, slice2)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
