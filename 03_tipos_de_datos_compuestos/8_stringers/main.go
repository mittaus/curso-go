package main

import "fmt"

//Persona x
type Persona struct {
	Name string
	Age  int
}

func (p Persona) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Persona{"Arthur Dent", 42}
	z := Persona{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
