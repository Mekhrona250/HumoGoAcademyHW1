package main
import "fmt"
func main() {
	var dlinaMassiva int
	var res int
	fmt.Scan(&dlinaMassiva)

	array := make([]int, dlinaMassiva)

	var q bool
	
	for i:=0; i < dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		

	}

	for i := 1; i < dlinaMassiva-1; i++ {
		if array[i] > array[i+1] && array[i] > array[i-1]  {
			res++
			q=true
			
		} 
	}

	if q {
		fmt.Println(res)

	} else {
		fmt.Println(0)
	}
		
	} 
    

	