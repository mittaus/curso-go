package main

import (
	"fmt"
	"reflect"
)

//Figura geometrica
type Figura struct {
	Name  string
	Color string
	Year  int
}

func main() {
	// reflect.TypeOf() espera un interface{} como parámetro y nos devuelve un reflect.Type es decir conoceremos el tipo de la interfaz que le estamos pasando, en este caso una Figura.
	// Kind() a partir de un reflect.Type sabremos la clase de tipo que es, en este caso struct.
	// reflect.ValueOf() dado un interface{} podremos averiguar su valor.

	g := Figura{Name: "Circulo", Color: "Rojo", Year: 20}
	rType := reflect.TypeOf(g)
	rKind := rType.Kind()
	rValue := reflect.ValueOf(g)

	fmt.Println(rType) //Nombre del tipo de variable de declaración
	fmt.Println(rKind) //tipo de dato objeto go
	fmt.Println(rValue)

	// name := rValue.Name //error por falta de hacer cast
	// fmt.Println(name)

	// name := rValue.Interface().(Figura).Name
	// fmt.Println(name)

	gType := reflect.TypeOf(g)
	fmt.Printf("Type: %s Kind: %s\n", gType.Name(), gType.Kind())

	ptrG := &g
	ptrType := reflect.TypeOf(ptrG)
	fmt.Printf("Type: %s Kind: %s\n", ptrType.Elem(), ptrType.Kind())

	strNumFields := gType.NumField()
	for i := 0; i < strNumFields; i++ {
		field := gType.Field(i)
		fmt.Printf("Field Type: %s: %s Kind: %s\n", field.Name, field.Type.Name(), field.Type.Kind())
	}

}
