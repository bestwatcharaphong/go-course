package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {
	productName[0] = "Macbook"
	productName[1] = "Ipad"
	productName[2] = "AirPods"
	productName[3] = "Iphone"

	price := [4]float32{40000, 30000, 20000, 2000}
	fmt.Println(price)

	fmt.Println(productName)

}
