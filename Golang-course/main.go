package main

import (
	"fmt"

	"github.com/watcharapong/Hello/formater"
)

var numberInt, numberInt2 int = 100, 2000
var msg string = "Hello"

func main() {
	numberfloat := 25.4
	fmt.Println(numberfloat)
	fmt.Println(numberInt2)
	fmt.Println(numberInt)
	fmt.Println(msg)
	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberInt) + numberfloat)
	fmt.Println("Hello")
	formater.MyPrintln()
	fmt.Println("hello"+"world", 10)
}
