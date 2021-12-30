package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Errorf(err.Error())
	}

	sa := strings.Split(string(data), "\n")
	var input []int
	for i := range sa {
		res, err := strconv.Atoi(sa[i])
		if err != nil {
			fmt.Errorf(err.Error())
		}
		input = append(input, res)
	}

	increases := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}
	fmt.Printf("Increases: %d", increases)
}
