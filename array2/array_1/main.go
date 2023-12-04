package main
import "fmt"
func main() {
var n int
fmt.Scan(&n)

ar := make([]int, n)


for i := 0; i < n; i++ {
  fmt.Scan(&ar[i])
}
var min_element, max_element int = 100000, -100000
for i := 0; i < n; i++ {
	if ar[i] > 0 && ar[i] < min_element {
		min_element= ar[i]
	}
	if ar[i] < 0 && ar[i] > max_element {
		max_element= ar[i]
	}
}
for i := 0; i<len(ar); i++ {
	if ar[i] == max_element || ar[i]==min_element {
		if i==0 {
			ar = ar[1:]
		} else if i < len(ar)-1 {
			ar = append(ar[:i], ar[i+1:]...)
		} else {
			ar = ar[:len(ar)-1]
		}
		i--
	}
}



fmt.Println(ar)
}