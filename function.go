package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello World")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plusThreeValue(value1 int, value2 int, value3 int) int {
	return value1 + value2 + value3
}

func main() {

	hello()
	// result := plus(1, 32)
	result := plusThreeValue(1, 2, 3)
	fmt.Println(result)
}
