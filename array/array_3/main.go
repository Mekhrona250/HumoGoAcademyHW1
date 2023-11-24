package main
import "fmt"
func main() {
	var dlinaMassiva int
	var max_element int
	fmt.Scan(&dlinaMassiva)

	array := make([]int, dlinaMassiva)

	
	for i:=0; i < dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		

	}
	max_element= array[0]
	for i := 1; i < dlinaMassiva; i++ {
		if array[i] > max_element {
			max_element= array[i]
		}
		
	}
    fmt.Println(max_element)

	}