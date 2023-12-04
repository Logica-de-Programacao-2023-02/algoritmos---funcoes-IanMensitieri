package main

import (
	"errors"
	"fmt"
	"unicode"
)

func stringsComMaiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice não pode estar vazio")
	}

	var result string
	for _, s := range slice {
		if len(s) > 0 && unicode.IsUpper([]rune(s)[0]) {
			result += s + " "
		}
	}

	return result[:len(result)-1], nil
}

func main() {
	sliceExemplo := []string{"Gato", "Cachorro", "elefante", "Tigre"}

	resultado, err := stringsComMaiuscula(sliceExemplo)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(resultado)
}
