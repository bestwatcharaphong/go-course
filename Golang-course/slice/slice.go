package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "html", "js", "css")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
