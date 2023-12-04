package main

import "fmt"

func main()  {
	var n int
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0;i < n; i++{
		fmt.Scan(&ar[i])
	}
	for i := 0; i < n / 2; i++{
		fmt.Print(ar[i], " ", ar[n - i - 1], " ")
	}
	if n % 2 == 1 {
		fmt.Print(ar[n / 2])
	}

	
}