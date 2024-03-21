package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Si el valor del canal está en dos, quiere decir que serán dos los procesos
	que se lancen, si se configura a 5, entonces serán 5 los procesos que se lancen
	entonces si es posible configurar cuantos procesos se quieren lanzar simultáneamente

	Cuando termina alguno de los procesos en ser ejecutado, entonces ahí si puede entrar otro proceso
*/

func main4() {
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething2(i, &wg, c)
	}

	wg.Wait()
}

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)
	<-c
}
