package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Tipo de Dato String
	// Los Strings son una secuencia de Bytes (Slice de Bytes)
	// Son Indexables
	// Son Inmutables

	var cadena string
	fmt.Println(cadena)
	cadena = "Hola Mundo"
	fmt.Println(cadena)
	fmt.Println(len(cadena) - 1)

	fmt.Println(cadena[9])
	//
	fmt.Println(cadena[:])
	//
	//cadena[2] = "l"
	//
	//
	cadena = cadena + " nuevamente"
	fmt.Println(cadena)
	//
	cadena += " siiii"
	fmt.Println(cadena)
	//
	cadena = `
<html>
	<head>
		<meta charset="utf-8">
			<title></title>
		</head>
		<body>

		</body>
	</html>
	`
	fmt.Println(cadena)

	cadena = "Hola Mundo \\\t\"25\""
	fmt.Println(cadena)

	edad := 29

	cadena = "La edad es " + strconv.Itoa(edad)

	fmt.Println("Edad", edad)

	//concatenar cadenas
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
	fmt.Println(b)

	/**************************  Inicio de interpolación de cadenas  **************************/

	//Interpolación de cadenas: Usando Print y Println del paquete fmt
	fmt.Println("Hola mundo", "¿Cómo estas?", true, 2)
	fmt.Print("Nuevo mensaje sin salto de línea", false, 2, 3.14)
	fmt.Print("\n")

	nombre := "Juan Carlos"
	miedad := 26
	activo := true

	fmt.Println("Hola", nombre, "con una edad", miedad, "y un status", activo)

	//Interpolación de cadenas: Usando Printf del paquete fmt
	// Se reemplaza conforme al orden de los siguientes argumentos.
	// El primer placeholder será reemplazado con el segundo argumento, el segundo placeholder con el tercer argumento, y así sucesivamente.
	miNombre2 := "Joel"
	fmt.Printf("Hola %v", miNombre2)

	fmt.Printf("Hola %v, con una edad de %v y un estatus %v ", "Edgar", 10, true)

	//Placeholder: %v (default format)
	fmt.Printf("String: %v \n", "Edgar")
	fmt.Printf("Entero: %v \n", 220)
	fmt.Printf("Flotante: %v \n", 3.1012)
	fmt.Printf("Booleano: %v \n", true && false)

	//Placeholder: %T (Type format)
	fmt.Printf("Tipo de dato: %v \n", "Edgar")
	fmt.Printf("Tipo de dato: %v \n", 220)
	fmt.Printf("Tipo de dato: %v \n", 3.1012)
	fmt.Printf("Tipo de dato: %v \n", true && false)

	//Placeholder: %t (Boolean Format)
	fmt.Printf("%t \n", true)
	fmt.Printf("%t \n", false)

	//Placeholder: %d (Base10 Format)
	fmt.Printf("%d \n", 124)
	fmt.Printf("%d \n", 0x75)

	//Placeholder: %b (Binary Format)
	fmt.Printf("%b \n", 124)
	fmt.Printf("%b \n", 0x75)

	//Placeholder: %f (Decimal Format)
	fmt.Printf("%f \n", 3.141592653589793238462643383)

	//Placeholder: %c (Character Format)
	mensaje1 := "Hola mundo"
	firstCharacter := mensaje1[0]

	fmt.Printf("%c \n", firstCharacter)
	fmt.Printf("%c \n", '🎉')

	//Placeholder: %s (String Format)
	mensaje2 := "Hola mundo"

	fmt.Printf("%s \n", mensaje2)
	fmt.Printf("%s \n", "Hola mundo con Emoji 🎉")

	//Placeholder: %p (Pointer address Format)
	mensaje3 := "Hola mundo"
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("%p \n", &mensaje3)
	fmt.Printf("%p \n", slice)

	//Interpolación de cadenas: Usando Sprintf del paquete fmt
	// A diferencia de la función Printf, Sprintf retorna el string generado a partir de la interpolación.
	otroNombre := "Eduardo"
	otraEdad := 30
	otroEstado := true

	textoInicial := "Hola %s con una edad %d y un status %t"

	mensajeFinal := fmt.Sprintf(textoInicial, otroNombre, otraEdad, otroEstado)
	fmt.Println(mensajeFinal)

}
