package main

import (
	"calc/solve"
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(errors.New("no arguments provided"))
		return
	}
	if len(os.Args) > 2 {
		fmt.Println(errors.New("too many arguments provided"))
		return
	}
	input := os.Args[1]
	output, err := solve.Solve(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(output)
}
