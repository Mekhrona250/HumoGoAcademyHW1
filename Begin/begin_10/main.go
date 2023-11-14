package main
import "fmt"
func main() {
	var a,b int
	fmt.Println("Введите а")
	fmt.Scan(&a)

	fmt.Println("Введите b")
	fmt.Scan(&b)

	fmt.Println((a*a)+(b*b))
	fmt.Println((a*a)-(b*b))
	fmt.Println((a*a)*(b*b))
	fmt.Println((a*a)/(b*b))
}