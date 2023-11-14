package main

import "fmt"
func main() {
	var a, b int

	var result bool
	
	fmt.Println("Введите числа")
	fmt.Scan(&a, &b)
	
	result = a>2 && b<=3

	fmt.Println(result)
}