package main

//Channels

import (
	"fmt"

	"time"
)

func imprimirPing(c chan string) {
	var contador int
	for {
		//Recibiendo valores a través del canal
		contador++
		fmt.Println(<-c, " ", contador)
		time.Sleep(time.Second * 1)
	}
}

func enviarPing(c chan string) {

	for {
		//Enviando valores a través del canal
		c <- "Ping"
	}
}
func main() {

	// I. A send to a nil channel blocks forever
	var c1 chan string
	c1 <- "Hello, World!"
	// fatal error: all goroutines are asleep - deadlock!

	// II. A receive from a nil channel blocks forever
	var c2 chan string
	fmt.Println(<-c2)
	// fatal error: all goroutines are asleep - deadlock!

	// III. A send to a closed channel panics
	var c3 = make(chan string, 1)
	c3 <- "Hello, World!"
	close(c3)

	// Creamos un canal hola mundoo
	ch := make(chan string)

	//Llamamos las funciones como Gorutines
	go enviarPing(ch)
	go imprimirPing(ch)

	//Escaneamos la entrada de datos para que no finalice la gorutine “main”
	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin...")

}
