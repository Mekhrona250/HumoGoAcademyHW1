package main
 
import "fmt"
func main() {
	var a,b float64
	

	fmt.Println("Введите число")
	fmt.Scan(&a,&b)

    if a>b {
		a,b= b,a
	}
	
	fmt.Println(b,a)
}
