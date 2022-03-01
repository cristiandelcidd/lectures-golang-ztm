package main

import (
	"fmt"
)

func double(num int) int {
	return num * 2
}

func add(number1, number2 int) int {
	return number1 + number2
}

func greet() {
	fmt.Println("Hello")
}

func main() {
	doubleNumber := double(3)
	fmt.Println(doubleNumber)

	result := add(1, 2)
	fmt.Println(result)

	result2 := add(double(4), 200)
	fmt.Println(result2)

	greet()
}
