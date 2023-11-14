package main
import(
	"fmt"
	"math"
)
func main() {
	var a,b float64
	fmt.Println("Введите неотрицательное число а")
	fmt.Scan(&a)

	fmt.Println("Введите неотрицательное число b")
	fmt.Scan(&b)
	
    fmt.Println(math.Sqrt(a*b))
	
}