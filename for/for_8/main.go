package main

import "fmt"

func main() {


	var a,b int
	fmt.Scan(&a,&b)

    var result int=1
	
for i :=a; i<=b; i++ {
	
    result *=i
}	

fmt.Println(result)
    
 


}