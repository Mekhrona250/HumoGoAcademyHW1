package main

import (
	"fmt"
	
)

func main() {


	var n int
	fmt.Scan(&n)
    
	var result float64 =1
	var a float64=1.1

	
for i :=1; i<=n; i++ {
    
	result *= a
	a +=0.1
}	
	
fmt.Println(result)
    
 


}