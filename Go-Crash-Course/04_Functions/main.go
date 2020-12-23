package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(greeting("Gran"))
	fmt.Println(getSum(3, 4))
}

func greeting(name string) string {
	return "Hello, " + name
}

func getSum(number1 int, number2 int) int { // can also be done by ... (number1, number2 int) int ...
	return number1 + number2
}
