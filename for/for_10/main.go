package main

import "fmt"

func main() {


	var n float64
	fmt.Scan(&n)

	var result float64 =0

	
for i :=1.0; i<=n; i++ {
    
	result +=1 / i 
}	
	
fmt.Println(result)
    
 


}