package main 
import "fmt"

func main() {
	var n,k int
    fmt.Scan(&n)
    k=1

   

	for k*k <= n {
		k++
    }
		
	fmt.Println(k)
       	
}

// 10 -> 4   2*2=4 3*3=9 4*4=16
// 1 - 1
// 2 - 4
// 3 - 9