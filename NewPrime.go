package main

import "fmt"
import "os"

func main() {
	a := 11
	b:=0
	if a==1||a==0{
		fmt.Println("Invalid Input")
		os.Exit(0)
	}
	for i := 2; i <= (a / i); i++ { // 	

		for a%i == 0 { // if i is a factor
			a = a / i // divide a by i else we wont get a prime number
			fmt.Println(a, " x ", i)
			b=i
		}
	}

	if a > 1 {
		fmt.Println("largest prime factor: ", a)
	}else{
		fmt.Println("Largest prime factor: ", b)
	}
}
