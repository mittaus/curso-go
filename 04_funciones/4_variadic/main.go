package funcion4

import "fmt"

func sumar(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}

func print(cadena string, cadenas ...string) {
	for _, c := range cadenas {
		cadena += " "
		cadena += c
	}
	fmt.Println(cadena)
}

func main() {

	//variadic function

	fmt.Println(sumar(1, 2, 3, 4))
	fmt.Println(sumar(10, 20, 30, 40))
	fmt.Println(sumar(10, 20))
	fmt.Println(sumar(20))
	fmt.Println(sumar())

	numeros := []int{
		10,
		20,
		30,
	}
	fmt.Println(sumar(numeros...))

	print("Hola", "mundo", "desde", "Go(lang)")

}
