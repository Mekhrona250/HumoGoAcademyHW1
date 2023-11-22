package main 
import "fmt"

func main() {
	var n int

	fmt.Scan(&n)



	for n>1 && n%3==0 {
		n/=3
    }
		if n==1{
			
			fmt.Println(true)
           
		} else {
			fmt.Println(false)
		}
	
       	
}