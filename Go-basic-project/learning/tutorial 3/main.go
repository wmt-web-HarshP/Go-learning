package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "HELLO WOrld"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 2
	var reminder, result, err = intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case reminder == 0:
		fmt.Println("The result of the integer division is %v", result)
	default:
		fmt.Println("The result of the integer division is %v with remainder %v", result, reminder)
	}
	switch reminder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	fmt.Println(result, reminder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Can't devided by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var reminder int = numerator % denominator
	return result, reminder, err
}
