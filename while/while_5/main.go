package main 
import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

    var result int

	for n>1 && n%2==0 {
		n/=2
		result++
    }
		
	fmt.Println(result)
       	
}