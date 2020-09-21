package main

import (
	"fmt"
	"reflect"
	"time"
)

//Persona es una persona
type Persona struct {
	Nombre string
	Edad   int
}

//Doc documento
type Doc struct {
	Persona
	Created time.Time
}

func main() {
	persona1 := Persona{"Joel", 10}
	persona2 := Persona{"Edgar", 10}
	persona3 := Persona{"Joel", 10}

	fmt.Print(persona1, "\n")
	fmt.Print(persona2, "\n")
	fmt.Print(persona1.Edad == persona2.Edad, "\n") //true

	//Usando reflect
	fmt.Print(reflect.DeepEqual(persona1, persona3), "\n") //true
	//Usando go-cmp de google: **Solo para escribir test
	//fmt.Print(cmp.Equal(persona1, persona3), "\n") //true

	d1 := Doc{
		Persona: Persona{"Bob", 21},
		Created: time.Now(),
	}
	time.Sleep(time.Millisecond)
	d2 := Doc{
		Persona: Persona{"Bob", 21},
		Created: time.Now(),
	}

	fmt.Println(d1 == d2)                 // false
	fmt.Println(d1.Persona == d2.Persona) // true

	fecha1 := time.Now()
	fecha2 := fecha1
	punteroFecha1 := &fecha1
	punteroFecha2 := &fecha2
	fmt.Print(fecha1, "\n")
	fmt.Print(fecha2, "\n")
	fmt.Print(&fecha1, "\n")
	fmt.Print(&fecha2, "\n")
	fmt.Print(*punteroFecha1, "\n")
	fmt.Print(*punteroFecha2, "\n")

	d3 := Doc{
		Persona: Persona{"Bob", 21},
		Created: fecha1,
	}
	time.Sleep(time.Millisecond)
	d4 := Doc{
		Persona: Persona{"Bob", 21},
		Created: fecha2,
	}
	fmt.Println(d3 == d4) // true
}
