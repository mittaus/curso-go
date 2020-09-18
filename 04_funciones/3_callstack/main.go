package main

import "fmt"

func funcion1() {
	fmt.Println("Ingreso a function1")
	funcion2()
	fmt.Println("Salida de funcion1")
}
func funcion2() {
	fmt.Println("Ingreso a function2")
	funcion3()
	fmt.Println("Salida de funcion2")
}
func funcion3() {
	fmt.Println("Ingreso a function3")

	fmt.Println("Salida de funcion3")
}

func main() {
	//call stack

	fmt.Println("Ingreso a main")
	funcion1()
	fmt.Println("Salida de main")

}
