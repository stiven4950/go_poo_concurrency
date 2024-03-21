package main

import "fmt"

func NewServerUser(name string, ram float64) *server {
	return &server{
		name:         name,
		ram_quantity: ram,
	}
}

type server struct {
	name         string
	ram_quantity float64
}

func main1() {
	s := NewServerUser("AWS", 13.2)
	fmt.Printf("%v\n", s)
}
