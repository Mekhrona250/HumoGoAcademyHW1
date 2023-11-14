package main
import "fmt"
func main() {
	var a,b int
	fmt.Println("Введите сторону а")
	fmt.Scan(&a)

	fmt.Println("Введите сторону b")
	fmt.Scan(&b)
	fmt.Println(2*(a+b))
	
}