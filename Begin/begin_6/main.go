package main
import "fmt"
func main() {
	var a,b,c int
	fmt.Println("Введите длину ребра а")
	fmt.Scan(&a)

	fmt.Println("Введите длину ребра b")
	fmt.Scan(&b)

	fmt.Println("Введите длину ребра c")
	fmt.Scan(&c)

	fmt.Println(a*b*c)
	fmt.Println(2*(a*b +b*c+a*c))
}