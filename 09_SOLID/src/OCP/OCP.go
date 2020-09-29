// go run main.go
// Example open/closed principle
// strategy pattern

package OCP

import "fmt"

// Calculer ...
type Calculer interface {
	Execute(rune, rune) rune
}

// Add ...
type Add struct{}

// Minus ...
type Minus struct{}

// Multiplicacion ...
type Multiplicacion struct{}

// Execute ...
func (a Add) Execute(numA, numB rune) rune {
	return numA + numB
}

// Execute ...
func (m Multiplicacion) Execute(numA, numB rune) rune {
	return numA * numB
}

// Execute ...
func (m Minus) Execute(numA, numB rune) rune {
	return numA - numB
}

// Calcul ...
type Calcul struct {
	c Calculer
}

// Calculate ...
func (c Calcul) Calculate(numA, numB rune) rune {
	return c.c.Execute(numA, numB)
}

func main() {
	fmt.Println("Open - closed Principle (OCP)")
	a := Calcul{c: Add{}}
	b := Calcul{c: Minus{}}
	c := Calcul{c: Multiplicacion{}}

	fmt.Println(a.Calculate(10, 5))
	fmt.Println(b.Calculate(10, 5))
	fmt.Println(c.Calculate(10, 5))

	fmt.Println(".....................................")
}
