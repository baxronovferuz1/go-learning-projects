package main

import "fmt"

func main() {
	var choice int
	var num1 float64
	var num2 float64

	fmt.Println("=== Calculator CLI ===")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	fmt.Print("Choice operation: ")
	fmt.Scan(&choice)

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	switch choice {
	case 1:
		fmt.Println("Result:", num1+num2)
	case 2:
		fmt.Println("Result:", num1-num2)
	case 3:
		fmt.Println("Result:", num1*num2)
	case 4:
		fmt.Println("Result:", num1/num2)
	default:
		fmt.Println("Invalid choices")
	}
}
