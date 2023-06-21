package main

import "fmt"

var product = make(map[string]float32)

func main() {
	product["Macbook"] = 400000
	product["ipone"] = 30000
	product["Ipad"] = 20000
	fmt.Println("product =", product)

	//delete
	delete(product, "ipone")
	fmt.Println(product)

	product["iphone"] = 20
	fmt.Println("product =", product)

	value1 := product["iphone"]

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("CousesName", courseName)

	fmt.Println(value1)
}
