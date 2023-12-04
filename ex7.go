package main

import (
	"errors"
	"fmt"
)

func applyFunctionToSlice(slice []int, f func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice is empty")
	}

	result := make([]int, len(slice))
	for i, v := range slice
