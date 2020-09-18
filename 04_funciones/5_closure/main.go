package main

import "fmt"

func print(cadena string) {
	fmt.Println(cadena)
}

func print2(cadena string) {
	fmt.Println(cadena)
}
func print3(cadena1, cadena2 string) {
	fmt.Println(cadena1 + cadena2)
}

// func print4(fprint func(string)) {
// 	fprint("Hola mundo desde Print4")
//
// }

func incremementar() func() int {
	i := 0
	return func() (r int) {
		r = i
		i += 2
		return
	}
}

func incremento() {
	i := 0
	i++
	fmt.Println(i)
}

func main() {

	//closure

	cadena := "Hola mundo"

	imprimir := print

	imprimir(cadena)

	imprimir2 := func() {
		fmt.Println(cadena)
	}

	imprimir2()

	imprimir = print2
	imprimir("Hola mundo 2")

	//imprimir = print3 //error por la firma

	fmt.Printf("Funcion print1: %T\n", print)
	fmt.Printf("Funcion print2: %T\n", print2)
	fmt.Printf("Funcion print3: %T\n", print3)

	//print4(print) //Se ejecuta sin problemas

	//Las funciones son comparables con (nil: valor no inicializado)
	var fb func()
	if fb == nil {
		fmt.Println("fb es igual a nil")
	}

	inc := incremementar()

	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())

	incremento()
	incremento()
	incremento()

}
