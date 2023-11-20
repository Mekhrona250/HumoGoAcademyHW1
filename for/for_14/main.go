package main

import "fmt"

func main() {


	var n int
	fmt.Scan(&n)
    
	var result int
	

	
for i :=1; i<=(2*n-1); i+=2 {
    
	result += i
	fmt.Println(result)	
}
    
 


}