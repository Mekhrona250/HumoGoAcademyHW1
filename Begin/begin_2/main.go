package main
import "fmt"
func main() {
	var a int

	fmt.Println("Привет, введите сторону квадрата")
	
	fmt.Scan(&a)
	
	fmt.Println(a*a)
}