package main

import "fmt"
func main() {
	var a float64

	fmt.Println("Введите число: ")
	fmt.Scan(&a)
	
	fmt.Println(a/1024)
}