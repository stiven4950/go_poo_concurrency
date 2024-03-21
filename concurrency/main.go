package main

import "fmt"

func main2() {
	// c := make(chan int) // Unbuffered
	c := make(chan int, 2) // Buffered

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
}
