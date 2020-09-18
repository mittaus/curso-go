// package main

// import "fmt"

// var var1 = "Este es el nivel 1"

// func main() {
// 	var var2 = "Este es el nivel 2"
// 	{
// 		var var3 = "Este es el nivel 3"
// 	}
// 	fmt.Println(var1)
// 	fmt.Println(var2)
// 	fmt.Println(var3)
// }

package main

import "fmt"

var razaThanos = "superhumanos"

func main() {
	//Nombrar Variables
	numero := 25
	_nombre := "Alejandro"
	numero2 := 54
	nombreUsuario := "admin"
	Numero := 46
	fmt.Println(numero, numero2, _nombre,
		nombreUsuario, Numero)
	var (
		villano                  = "Thanos"
		superHeroe1, superHeroe2 = "Iron man", "Hulk"
	)
	var numero3 = 66
	fmt.Println(villano, superHeroe1, superHeroe2, numero3)

	//Scope

	fmt.Println("El villano es: " + razaThanos)
	imprimir()

}

func imprimir() {
	fmt.Println("La raza de Thanos es: " + razaThanos)
}
