package main

import "fmt"
func main() {
	var a,b int
	

	fmt.Println("Введите число")
	fmt.Scan(&a,&b)

	if a<b {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}

}