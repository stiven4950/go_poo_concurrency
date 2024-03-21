package main

import "fmt"

/*
	Cuando se trabaja con channels existe la gran posibilidad
	de crear un deadlock si no somos cuidadosos con su utilización,
	una forma de motigar parte de este riesgo es definiendo canales
	de lectura o escritura, pero no ambos.
*/

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}

	close(out)
}

func Print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main5() {
	c := make(chan int)
	out := make(chan int)

	go Generator(c)
	go Double(c, out)
	Print(out)
}
