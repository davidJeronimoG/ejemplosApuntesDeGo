package main

import "fmt"

func leerEscribir(c chan string) {
	c <- "Mazda"
	fmt.Println(<-c)
}

func soloLectura(c <-chan string) {
	fmt.Println(<-c)
}

func soloEscribir(c chan<- string) {
	c <- "Kia"
}

func main1() {

	autos := make(chan string, 1)
	leerEscribir(autos)
	autos <- "Toyota"
	soloLectura(autos)
	soloEscribir(autos)
	fmt.Println(<-autos)
}
