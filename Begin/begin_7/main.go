package main
import "fmt"
func main() {
	var r float64
	fmt.Println("Введите радиус")
	fmt.Scan(&r)
	fmt.Println(2*3.14*r)
	fmt.Println(3.14*r*r)
}