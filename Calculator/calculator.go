package main

import ("fmt")

func GetInput() (string, int, int) {
	var operationType string
	var number1 int
	var number2 int

	fmt.Println("Choose an operation (+, -, *, /): ")
	fmt.Scan(&operationType)

	fmt.Println(": ")
	fmt.Scan(&number1)

	fmt.Println(": ")
	fmt.Scan(&number2)
	return operationType, number1, number2
}

func CalculateResult(operation string, n1 int, n2 int) int {
	
	
	return 
}

func main(){
	CalculateResult(GetInput())
}