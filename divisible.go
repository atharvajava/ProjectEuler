package main

import "fmt"

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println(gcd(i, j))
		}
	}
	fmt.Println(gcd(8, 10))
}
