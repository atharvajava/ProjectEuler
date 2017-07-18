package main

import (
	"fmt"
)

func isPrime(value int) bool {
	for i := 2; i <= value/2; i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func prime(x int) int {
	count := 0
	val := 0
	for i := 1; count < x; i++ {
		if isPrime(i) {
			count++
			val = i
		}
	}
	return val
}

func main() {
	fmt.Print(prime(10001))
}
