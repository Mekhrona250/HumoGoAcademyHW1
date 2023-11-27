package main
import "fmt"
func main() {
	var dlinaMassiva int
	fmt.Scan(&dlinaMassiva)
	var k int=0

	array := make([]int, dlinaMassiva)

	
	for i:=0; i<dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		
	}
	
	var x int
	fmt.Scan(&x)



	for i := 0; i<dlinaMassiva; i++ {
		if array[i]==x {
			k++
			fmt.Println(i+1)
		
		} 

	
	}



}