package main

import "fmt"

func reverse(a int) int {
	var reverse int
	for a != 0 {
		reverse = reverse * 10
		reverse = (reverse + a%10)
		a = a / 10
	}
	//fmt.Println(reverse)
	return reverse
}
func main() {
	pal := 0
	re := 0
	for i := 900; i <= 999; i++ {
		for j := 900; j <= 999; j++ {
			pal = j * i
			re = reverse(pal)
			if pal == re {
				fmt.Println("This is pal", j, "X", i, "=", pal)
			}
		}
	}
	reverse(121)
}
