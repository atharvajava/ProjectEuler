package main

import (
	"fmt"
	"math"
)

//IsPrimeSqrt checks if the value is prime or not
func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

//LargestPrime Prints the largest Prime factor
func LargestPrime(value int) int {
	var a int
	for i := 2; i < (value / i); i++ {
		if value%i == 0 {
			if IsPrimeSqrt(i) {
				fmt.Println(i)

			}
		}
	}
	return a
}

func main() {
	fmt.Println(LargestPrime(13195))
}
