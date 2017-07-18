package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	sum = sum * sum
	for i := 1; i <= 100; i++ {
		sum -= i * i
	}
	fmt.Println(sum)
}
