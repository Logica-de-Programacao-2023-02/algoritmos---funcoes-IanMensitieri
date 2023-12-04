package main

import (
	"errors"
	"fmt"
)

func filterEvenNumbers(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice is empty")
	}

	var evenNumbers []int
	for _, num := range slice {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}

	return evenNumbers, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result, err := filterEvenNumbers(slice)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Even Numbers:", result)
	}
}
