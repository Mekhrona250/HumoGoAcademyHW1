package main 
import "fmt"

func main() {
	var n,k int

	fmt.Scan(&n,&k)

	var result int

	for n>=k {
		n -=k
		result++

    
	}

fmt.Println(result)
fmt.Println(n)
}