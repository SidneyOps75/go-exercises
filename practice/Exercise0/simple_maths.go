package main

import (
	"fmt"
	"math"
)

func sum(x, y float64) float64 {
	return x + y
}

func sqrtOfSum(x, y float64) float64 {
	return math.Sqrt(sum(x, y))
}

func main() {
	var x float64
	var y float64
	fmt.Println("Input the first number: ")
	fmt.Scanln(&x)
	fmt.Println("Input the second number: ")
	fmt.Scanln(&y)
	fmt.Println("The sum: ", sum(x, y))
	fmt.Println("The square of the sum: ", sqrtOfSum(x, y))
}
