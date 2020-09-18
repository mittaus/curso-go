package main

import (
	"errors"
	"fmt"
)

//Errores

var (
	//ErrorUsuarioNoValido ...
	ErrorUsuarioNoValido = errors.New("El usuario no es valido")
	//ErrorUsuarioEnProceso ...
	ErrorUsuarioEnProceso = errors.New("Usuario en proceso de registro")
	//ErrorPorDefecto ...
	ErrorPorDefecto = errors.New("Algo paso")
)

func baneado(usuario string) (err error) {
	estaBaneado := false
	switch usuario {
	case "miguel":
		estaBaneado = true
	case "carlos":
		estaBaneado = false
	case "juan":
		return fmt.Errorf("El usuario no es valido")
	case "pedro":
		return fmt.Errorf("Usuario en proceso de registro")
	default:
		return fmt.Errorf("Algo paso")
	}

	if !estaBaneado {
		fmt.Printf("El usuario %s no esta baneado \n", usuario)
	} else {
		fmt.Printf("El usuario %s esta baneado \n", usuario)
	}

	return nil

}

// func baneado(usuario string) (err error) {
// 	estaBaneado := false
// 	switch usuario {
// 	case "miguel":
// 		estaBaneado = true
// 	case "carlos":
// 		estaBaneado = false
// 	case "juan":
// 		return ErrorUsuarioNoValido
// 	case "pedro":
// 		return ErrorUsuarioEnProceso
// 	default:
// 		return ErrorPorDefecto
// 	}

// 	if !estaBaneado {
// 		fmt.Printf("El usuario %s no esta baneado \n", usuario)
// 	} else {
// 		fmt.Printf("El usuario %s esta baneado \n", usuario)
// 	}

// 	return nil
// }

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	err := baneado("miguel")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = baneado("juan")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = baneado("carlos")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = baneado("pedro")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = baneado("pololo")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// err := baneado("miguel")
	// checkError(err)

	// err = baneado("juan")
	// checkError(err)

	// err = baneado("carlos")
	// checkError(err)

	// err = baneado("pedro")
	// checkError(err)

	// err = baneado("pololo")
	// checkError(err)

}
