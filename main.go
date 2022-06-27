package main

import "fmt"

func main() {

	miCanal := make(chan string, 5)

	// go func() {
	// 	miCanal <- "Hola"
	// }()

	miCanal <- "primera linea"
	miCanal <- "segunda linea"
	miCanal <- "tercera linea"

	close(miCanal)

	fmt.Println(<-miCanal)
fmt.Println("----------------------------------------------------")
	for v := range miCanal {
		fmt.Println(v)
	}
}
