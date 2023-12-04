package main

import (
	"fmt"
	"sort"
)

func segundoMaiorValor(slice []int) (int, error) {
	if len(slice) < 2 {
		return 0, fmt.Errorf("O slice deve conter pelo menos dois elementos")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	return slice[1], nil
}

func main() {
	valores := []int{10, 5, 8, 15, 7}

	segundoMaior, err := segundoMaiorValor(valores)

	if err == nil {
		fmt.Printf("O segundo maior valor é: %d\n", segundoMaior)
	} else {
		fmt.Println(err)
	}
}
