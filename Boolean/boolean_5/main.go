package main

import "fmt"
func main() {
	var a,b int
	
	var result bool

	fmt.Println("Введите числа")
	fmt.Scan(&a, &b)

	result =  a>=0 && b< -2

	fmt.Println(result)
}