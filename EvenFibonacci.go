package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 1
	var sum int
	var c int
	fmt.Println(a)
	fmt.Println(b)
	for i := 0; c < 4000000; i++ {
		c = a + b
		//fmt.Println(c)
		a = b
		b = c

		if c%2 == 0 {
			sum += c
			//fmt.Println(sum)
		}
	}
	fmt.Println(sum)
}
