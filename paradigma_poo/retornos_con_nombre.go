package main

import "fmt"

func getPows(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func main() {
	fmt.Println(getPows(2))
}
