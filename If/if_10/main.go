package main

import "fmt"
func main() {
	var a,b float64
	
    fmt.Println("Введите число")
	fmt.Scan(&a,&b)

    if a != b {
		a = a+b
		b = a
	} else {
		a = 0
		b = 0
	}

	fmt.Println(a, b)
	
}






