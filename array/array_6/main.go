package main
import "fmt"
func main() {
	var dlinaMassiva int
	fmt.Scan(&dlinaMassiva)

	array := make([]int, dlinaMassiva)

	
	for i:=0; i<dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		
	}
	
	var x int
	fmt.Scan(&x)

	var count int
	for i := 0; i<dlinaMassiva; i++ {
		if array[i]==x {
			count++ 
		}

	
	}
fmt.Println(count)
	}