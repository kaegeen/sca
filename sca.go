package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Prompt for user input
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	// Perform calculation based on operator
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Error: Invalid operator. Please use +, -, *, or /.")
	}
}
