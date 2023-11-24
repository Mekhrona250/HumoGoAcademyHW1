package main
import "fmt"
func main() {
	var dlinaMassiva int
	fmt.Scan(&dlinaMassiva)

	var res bool

	array := make([]int, dlinaMassiva)

	
	for i:=0; i < dlinaMassiva; i++ {
		fmt.Scan(&array[i])
		}

	for i :=0; i < dlinaMassiva-1; i++ {
		if (array[i]>0 && array[i+1]>0) || (array[i]<0 && array[i+1]<0) {
			res=true
			break
		}

	}
	
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	
	}


	

