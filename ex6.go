package main

import (
	"errors"
	"fmt"
	"strings"
)

func concatenarComVirgulas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}
	return strings.Join(slice, ","), nil
}

func main() {
	stringsParaConcatenar := []string{"Olá", "Mundo", "Gopher"}

	resultado, err := concatenarComVirgulas(stringsParaConcatenar)

	if err == nil {
		fmt.Printf("A concatenação com vírgulas é: %s\n", resultado)
	} else {
		fmt.Println(err)
	}
}
