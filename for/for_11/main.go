package main

import  "fmt"

func main() {


	var n int
	fmt.Scan(&n)

	var result int

	
for i :=0; i<=n; i++ {
    
	result += (n+i)*(n+i)
}	
	
fmt.Println(result)
    
 


}