package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	WAITGROUP
	Cuando se trabaja con Goroutines no siempre se planea que estas
	envíen datos a través de canales entre unas y otras, en estos casos
	se puede utilizar un Wait Group para alcanzar ese bloqueo de rutinas
	que es necesario.
*/

func main3() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething1(i, &wg)
	}

	wg.Wait()
}

func doSomething1(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
