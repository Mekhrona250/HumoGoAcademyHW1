package main
import "fmt"
func main() {
	var dlinaMassiva int
	fmt.Scan(&dlinaMassiva)

	array := make([]int, dlinaMassiva)

	
	for i:=0; i<dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		

	}
	for i := dlinaMassiva - 1; i>=0; i-- {
		fmt.Println(array[i])
	}


	}