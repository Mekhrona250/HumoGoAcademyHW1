package main 
import "fmt"

func main() {
	var n,k int
    fmt.Scan(&n)
    k=1

	var res int 

	for k < n {
		k*=3
		res++
    }
		
	fmt.Println(res-1)
       	
}