package main

import "fmt"
func main() {
	var a,b,c int
	

	fmt.Println("Введите число")
	fmt.Scan(&a,&b,&c)

	result := 0
    result2 := 0

	if a>0 {
		result += 1
	} else {
		result2 +=1
	}
	
	if b>0 {
		result +=1
		
	} else {
		result2 +=1
	}
	
	if c>0 {
	   result +=1
	} else {
		result2 +=1
	}


    fmt.Println(result) 
	fmt.Println(result2)
	
}