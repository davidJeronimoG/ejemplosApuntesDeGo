package main

import "fmt"

func leerEscribir(c chan string) {
	c <- "Mazda"
	fmt.Println(<-c)
}

func main() {

	autos := make(chan string, 3)

	leerEscribir(autos)

}
