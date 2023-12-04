package main

import (
	"fmt"
	"strings"
)

func concatenarStrings(slice []string) string {
	return strings.Join(slice, "")
}

func main() {
	stringsParaConcatenar := []string{"Olá, ", "Mundo", "!"}

	resultado := concatenarStrings(stringsParaConcatenar)

	fmt.Printf("A concatenação das strings é: %s\n", resultado)
}
