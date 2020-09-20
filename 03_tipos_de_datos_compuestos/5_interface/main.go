package main

import (
	"fmt"
)

//Persona ...
type Persona struct {
	nombre string
	email  string
	edad   int
}

//Nombre ...
func (p Persona) Nombre() string {
	return p.nombre
}

//Email ...
func (p Persona) Email() string {
	return p.email
}

//Moderador ...
type Moderador struct {
	Persona
	Foro string
}

//AbrirForo ...
func (m Moderador) AbrirForo() {
	fmt.Println("Abrir un Foro")
}

//CerrarForo ...
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar un Foro")
}

//Administrador ...
type Administrador struct {
	Persona
	Sección string
}

//CrearForo ...
func (a Administrador) CrearForo() {
	fmt.Println("Abrir un Foro")
}

//EliminarForo ...
func (a Administrador) EliminarForo() {
	fmt.Println("Cerrar un Foro")
}

//Presentarse ...
func Presentarse(p Persona) {
	fmt.Println("Nombre: ", p.Nombre())
	fmt.Println("Email: ", p.Email())
}

//PresentarseA ...
func PresentarseA(a Administrador) {
	fmt.Println("Nombre: ", a.Nombre())
	fmt.Println("Email: ", a.Email())
}

//PresentarseM ...
func PresentarseM(m Moderador) {
	fmt.Println("Nombre: ", m.Nombre())
	fmt.Println("Email: ", m.Email())
}

//Usuario ...
type Usuario interface {
	Nombre() string
	Email() string
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {

	alejandro := Persona{"Alejandro", "a@hmail.com", 29}
	juan := Moderador{Persona{"Juan", "j@hmail.com", 46}, "Juegos"}
	pedro := Administrador{Persona{"Pedro", "p@hmail.com", 25}, "PC"}

	Presentarse(alejandro)
	PresentarseM(juan)
	PresentarseA(pedro)

	// var i Usuario
	// i = alejandro
	// fmt.Println("i: ", i)
	// fmt.Println("i: ", i.Email())
	// fmt.Println("i: ", i.Nombre())

	// i = juan
	// fmt.Println("i: ", i)
	// fmt.Println("i: ", i.Email())
	// fmt.Println("i: ", i)

	// //Interfaz vacio
	// var i interface{}
	// describe(i)

	// i = 42
	// describe(i)

	// i = "Hola"
	// describe(i)

	// // Type assertions
	// s := i.(string)
	// fmt.Println(s)

	// s, ok := i.(string)
	// fmt.Println(s, ok)

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)

	// // Type switches
	// do(21)
	// do("hello")
	// do(true)
}
