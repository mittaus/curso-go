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
}
