package main

import "fmt"
func main() {
	var a,b int
	

	fmt.Println("Введите число")
	fmt.Scan(&a,&b)

    if a>b {
		a,b= b,a
	}
	
	fmt.Println(b,a)
}
// 5 2
// 3 4