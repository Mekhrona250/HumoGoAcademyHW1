package main
import "fmt"
func main() {
	var a int
	fmt.Println("Введите длину ребра куба а")
	fmt.Scan(&a)
	
	fmt.Println(a*a*a)
	fmt.Println(6*(a*a))
	
}
