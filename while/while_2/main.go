package main 
import "fmt"

func main() {
	var a,b int

	fmt.Scan(&a,&b)
    
	var result int

	for a>=b {
		a -=b
		result++
	}

fmt.Println(result)
}