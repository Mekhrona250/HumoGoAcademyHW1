package main
import "fmt"
func main() {
	var dlinaMassiva int
	var min_element int
	fmt.Scan(&dlinaMassiva)

	var res bool
	array := make([]int, dlinaMassiva)

	
	for i:=0; i < dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		

	}
	min_element= 1000000
	for i := 0; i < dlinaMassiva; i++ {
		if array[i] < min_element && array[i]%2 !=0 {
			min_element= array[i]
			res=true
		} 
		
	} 

	if res {
		fmt.Println(min_element)

	} else {
		fmt.Println(0)
	}

	
	}