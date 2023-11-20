package main

import "fmt"

func main() {
    var a float64
    fmt.Scan(&a)

	var n int
	fmt.Scan(&n)
    
	var result float64=1
	

	
for i :=1; i<=n; i++ {
    
	result *= a
}

fmt.Println(result)
    
 


}