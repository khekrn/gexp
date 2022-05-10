package main

import "fmt"

func main() {
	fmt.Println("Hello Go !!")
	fmt.Println("Sum of Array = ", sum(1, 2, 3, 4, 5))
}

func sum(arr ...int) int {
	res := 0
	for _, val := range arr {
		res += val
	}
	return res
}
