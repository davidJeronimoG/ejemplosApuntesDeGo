package main

import (
	"fmt"
	"time"
)

func main() {

	can1 := make(chan string)
	can2 := make(chan string)
	can3 := make(chan string)

	go caballo1(can1)
	go caballo2(can2)
	go caballo3(can3)

	var ganador string

	select {
	case ganador = <-can1:
	case ganador = <-can2:
	case ganador = <-can3:
	case <-time.After(time.Second * 5):
		ganador = "Todos se murieron"
	}

	fmt.Println("El ganador es:", ganador)
}

func caballo1(c chan<- string) {
	time.Sleep(time.Second * 6)
	c <- "Caballo 1"
}

func caballo2(c chan<- string) {
	time.Sleep(time.Second * 8)
	c <- "Caballo 2"
}

func caballo3(c chan<- string) {
	time.Sleep(time.Second * 10)
	c <- "Caballo 3"
}
