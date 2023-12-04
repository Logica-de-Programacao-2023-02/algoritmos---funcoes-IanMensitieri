package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarSlice(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice não pode estar vazio")
	}

	sortedSlice := make([]int, len(slice))
	copy(sortedSlice, slice)
	sort.Ints(sortedSlice)

	return sortedSlice, nil
}

func main() {
	sliceExemplo := []int{4, 2, 7, 1, 9}

	resultado, err := ordenarSlice(sliceExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
