package main
import "fmt"
func main() {
	var dlinaMassiva int
	var min_element int
	var max_element int
	fmt.Scan(&dlinaMassiva)

	array := make([]int, dlinaMassiva)

		min_element= 100
	    max_element= array[0]

	for i:=0; i < dlinaMassiva; i++ {
		fmt.Scan(&array[i])

	}


	for i := 1; i < dlinaMassiva; i++ {
		if array[i] < min_element {
			min_element= array[i]
		
		} else if array[i]> max_element {
			max_element = array[i]
			
		} 
	} 

	for i := 0; i < dlinaMassiva; i++ {

		if array[i] == max_element{
		
			fmt.Println(min_element)
			
		} else {
			fmt.Println(array[i])
		}
	}

	
	
	
	}





	

	