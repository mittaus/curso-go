package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Declarar un array y slice literal
	numeros := [6]int{2, 3, 5, 7, 11, 13}
	nombres := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	rType1 := reflect.TypeOf(numeros)
	rKind1 := rType1.Kind()

	rType2 := reflect.TypeOf(nombres)
	rKind2 := rType2.Kind()

	fmt.Println(rType1, rKind1)
	fmt.Println(rType2, rKind2)

	//Slices

	//Declarar Slices
	var j []int
	fmt.Println("Slice j: ", j)

	//Declarar Slices 2
	x := []int{1, 2, 3}
	fmt.Println("Slice x: ", x)

	//Declarar Slices usando Make, indicando la longitud
	y := make([]int, 5)
	fmt.Println("Slice y: ", y)
	fmt.Println("Longitud de y: ", len(y))
	fmt.Println("Capacidad de y: ", cap(y))

	//Declarar Slices usando Make indicando la longitud y la capacidad
	k := make([]int, 5, 10)
	fmt.Println("Slice k: ", k)
	fmt.Println("Longitud de k: ", len(k))
	fmt.Println("Capacidad de k: ", cap(k))

	// Definimos un array con 10 elementos de tipo int
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array ar: ", ar)

	// definimos dos slice de tipo []int
	var a, b []int
	fmt.Println("Slice a: ", a)
	fmt.Println("Slice b: ", b)

	// 'a' apunta del 3ro al 5to elemento en el array ar.
	a = ar[2:5]
	// ahora 'a' tiene los elementos ar[2],ar[3] y ar[4]
	fmt.Println("Slice a: ", a)

	// 'b' es otro slice del array ar
	b = ar[3:5]
	// Ahora 'b' tiene ar[3] y ar[4]
	fmt.Println("Slice b: ", b)

	b[0] = 25
	fmt.Println("Asignamos b[0] = 25")
	fmt.Println("Slice b: ", b)
	fmt.Println("Slice a: ", a)
	fmt.Println("Array ar: ", ar)
	fmt.Println("Cap(a)", cap(a))
	fmt.Println("Cap(b)", cap(b))

}
