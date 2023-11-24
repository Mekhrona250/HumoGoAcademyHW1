package main
import "fmt"
func main() {
	var dlinaMassiva int
	fmt.Scan(&dlinaMassiva)

	array := make([]int, dlinaMassiva)

	
	for i:=0; i<dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		

	}
	for i := 0; i < dlinaMassiva / 2; i++ {
		array[i], array[dlinaMassiva - i - 1]=array[dlinaMassiva - i -1], array[i]
	}
    fmt.Println(array)

	}