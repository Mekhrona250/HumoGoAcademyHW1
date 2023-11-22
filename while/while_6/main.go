package main 
import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

    var result float64 =1

	for n >=1.0 {
		result *= float64(n)
		n -= 2
    }
		
	fmt.Println(result)
       	
}

// 8*6*4*2
// 9*7*5*3*1