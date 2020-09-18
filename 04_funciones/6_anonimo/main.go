package main

import (
	"fmt"
	"strings"
)

func main() {

	//funciones anónimas

	cadena := "abcde"

	cadena = strings.Map(func(r rune) rune {
		return r + 1

	}, cadena)

	fmt.Println(cadena)

}
