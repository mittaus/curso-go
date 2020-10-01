package calc_test

import (
	"fmt"
	"testing"

	. "example.com/mittaus/go-testing-frameworks/calc"
)

func TestNameStandard(t *testing.T) {
	fmt.Println("************* STANDARD *****************")
}

func TestAddWithTestingPackage(t *testing.T) {

	operador1 := 1
	operador2 := 4

	result := Add(operador1, operador2)

	if result != 3 {
		t.Errorf("El resultado fue incorrecto, obtuvo: %d, se quería: %d.", result, 3)
	}
}

func TestSubtractWithTestingPackage(t *testing.T) {
	result := Subtract(5, 3)

	if result != 2 {
		t.Errorf("El resultado fue incorrecto, obtuvo: %d, se quería: %d.", result, 2)
	}
}

func TestMultiplyWithTestingPackage(t *testing.T) {
	result := Multiply(5, 6)

	if result != 30 {
		t.Errorf("El resultado fue incorrecto, obtuvo: %d, se quería: %d.", result, 30)
	}
}

func TestDivideWithTestingPackage(t *testing.T) {
	result := Divide(10, 2)

	if result != 5 {
		t.Errorf("El resultado fue incorrecto, obtuvo: %f, se quería: %f.", result, float64(5))
	}
}
