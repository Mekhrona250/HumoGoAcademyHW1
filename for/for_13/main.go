package main

import "fmt"


func main() {


	var n float64
	fmt.Scan(&n)
    
	var result float64 =0
	var a float64=1.0

	
for i :=1.0; i<=n; i++ {
    
	result += a*(1.0+0.1*i)
	a =a*(-1.0)

}
fmt.Println(result)
}