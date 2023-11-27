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
	var res bool
	for i := 0; i<dlinaMassiva; i++ {
		if array[i]==x {
			count++ 
			res=true
		}

	
	}

	if res {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")
	}






	}