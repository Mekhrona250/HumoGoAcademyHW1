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
	
	
	x = 1000															

 	for i:=1; i<dlinaMassiva; i++ {
		if x - array[i-1] > array[i] -x {
			
			fmt.Println(array[i-1])
			} 
		

	} 



}

