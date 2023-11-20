package main

import "fmt"
func main() {
	var a,b int
    fmt.Scan(&a,&b)
 
	
for i :=b-1; i>=a+1; i-- {
	fmt.Println(i)
    

}

fmt.Println(b-a-1)

}