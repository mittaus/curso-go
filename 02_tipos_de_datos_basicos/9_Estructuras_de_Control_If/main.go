package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v)
		fmt.Printf("%g >= %g\n", v, lim)
		// return v
	}

	// No se puede usar la variable aquí
	// fmt.Println(v)
	return lim
}

func main() {

	//Estructuras de Control: if

	// if 5 < 6 {
	// 	fmt.Println("5 es menor que 6")
	// }

	a := 6

	if a < 6 {
		fmt.Println("a es menor que 6")
	} else if a > 6 {
		fmt.Println("a es mayor que 6")
	} else {
		fmt.Println("a es igual 6")
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("Fin del programa")

}
