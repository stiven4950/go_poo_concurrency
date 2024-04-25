package main

import "fmt"

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("-> Started-- Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("--Ended-- Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{40, 41, 42, 43, 44, 45, 46, 47, 48}
	nWorkers := 10
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
