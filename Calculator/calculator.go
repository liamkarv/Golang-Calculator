package main

import ("fmt")

func GetInput() (string, int, int) {
	var operationType string
	var number1 int
	var number2 int

	fmt.Println("Choose an operation (+, -, *, /): ")
	fmt.Scan(&operationType)

	fmt.Println("Choose a number: ")
	fmt.Scan(&number1)

	fmt.Println("Choose another number: ")
	fmt.Scan(&number2)
	return operationType, number1, number2
}

func CalculateResult(operation string, n1 int, n2 int) int {
	var result int

	if operation == "+" {
		result = n1 + n2
	}
	
	if operation == "-" {
		result = n1 - n2
	}

	if operation == "*" {
		result = n1 * n2
	}

	if operation == "/" {
		result = n1 / n2
	}

	return result
}

func main(){
	fmt.Println("Calculation result:", CalculateResult(GetInput()))
}