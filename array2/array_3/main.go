package main
import "fmt"
func main() {
var n,k int
fmt.Scan(&n,&k)

ar := make([]int, n)


for i := 0; i < n; i++ {
  fmt.Scan(&ar[i])
}
for i :=0; i < n; i++ {
	fmt.Print(ar[i], " ")
	if (i+1)%k == 0 {
		fmt.Print("\n")
	}
}

}