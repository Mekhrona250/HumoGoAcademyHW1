package main

import "fmt"
func main() {
	var a int

	fmt.Println("Введите число: ")
	fmt.Scan(&a)
	
	fmt.Println(a/100)
	fmt.Println( a % 100 / 10)

}