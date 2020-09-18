package main

import "fmt"

func suma(n1 int, n2 int) int {
	return n1 + n2
}

func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

func main() {
	//firma de una función

	fmt.Printf("Funcion suma: %T\n", suma)
	fmt.Printf("Funcion resta: %T\n", resta)
}
