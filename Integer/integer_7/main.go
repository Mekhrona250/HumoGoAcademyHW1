package main

import "fmt"
func main() {
	var a int

	fmt.Println("Введите число: ")
	fmt.Scan(&a)
	
	fmt.Println((a/10)+(a%10))
	fmt.Println((a/10)*(a%10))
}