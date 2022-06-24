package main

import "fmt"

func main() {
	fmt.Println(Arithmetic(30, 3, "multiply"))
}

func Arithmetic(a int, b int, operator string) int {
	//your code here
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "divide":
		return a / b
	case "multiply":
		return a * b
	}
	return 0
}
