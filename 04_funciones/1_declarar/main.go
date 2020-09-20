package main

import (
	"fmt"
)

// versión 1
func imprimirNombre(nombre string) {
	fmt.Println("Fuera de Main")
	fmt.Println("El nombre es: ", nombre)
}

// versión 2
func suma(n1 int, n2 int) int {
	return n1 + n2
}

// versión 3
func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

// versión 3
func concatenar(n1, n2 int, n3 string) (r string) {
	r = fmt.Sprintf("%d.%v - %v", n1, n2, n3)
	return
}

func main() {

	//declara funciones
	imprimirNombre("Jose")
	fmt.Println("Dentro de main")

	fmt.Println(suma(25, 66))
	fmt.Println(resta(88, 66))
	resultado := concatenar(1, 2, "mundo")
	fmt.Println(resultado)
	//package math
	//func Sin(x float64) float64 // implemented in assembly language
}
