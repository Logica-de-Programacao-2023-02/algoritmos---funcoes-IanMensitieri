package main

import (
	"fmt"
	"strings"
	"unicode"
)

func contarVogais(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) && strings.ContainsRune("aeiouAEIOU", char) {
			count++
		}
	}
	return count
}

func main() {
	texto := "Hello, Gopher!"

	qtdVogais := contarVogais(texto)

	fmt.Printf("A quantidade de vogais na string é: %d\n", qtdVogais)
}
