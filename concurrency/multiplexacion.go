package main

import (
	"fmt"
	"time"
)

/*
	MULTIPLEXACIÓN

	Cuando una rutina se está comunicando con varios channels
	es muy útil usar la palabra reservada select para poder
	interactuar de una manera más ordenada con todos los mensajes
	que están siendo recibidos
*/

func main7() {

	c1 := make(chan int)
	c2 := make(chan int)

	d1 := time.Second * 3
	d2 := time.Second * 1

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case canalUno := <-c1:
			fmt.Println(canalUno)
		case canalDos := <-c2:
			fmt.Println(canalDos)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
