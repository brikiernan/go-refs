package main

import (
	"errors"
	"fmt"
)

func main() {

	printMe("Hey now!")

	var result, remainder, err = intDivision(10, 5)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the interer division is %v", result)
	} else {
		fmt.Printf("The result of the interer division is %v with remainder %v", result, remainder)
	}

	var result2, remainder2, err2 = intDivision(10, 5)
	switch {
	case err2 != nil:
		fmt.Printf(err2.Error())
	case remainder2 == 0:
		fmt.Printf("The result of the interer division is %v", result2)
	default:
		fmt.Printf("The result of the interer division is %v with remainder %v", result2, remainder2)

	}
}

func printMe(val string) {
	fmt.Println(val)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
