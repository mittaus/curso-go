package main

import (
	"fmt"
	"strings"
)

func main() {
	// Crear un tablero tic-tac-toe.
	tablero := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Los jugadores se turnan
	tablero[0][0] = "X"
	tablero[2][2] = "O"
	tablero[1][2] = "X"
	tablero[1][0] = "O"
	tablero[0][2] = "X"

	for i := 0; i < len(tablero); i++ {
		fmt.Printf("%s\n", strings.Join(tablero[i], " "))
	}
}
