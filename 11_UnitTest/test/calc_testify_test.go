package calc_test

import (
	"fmt"
	"testing"

	. "example.com/mittaus/go-testing-frameworks/calc"
	. "github.com/stretchr/testify/assert"
)

func TestNameTestify(t *testing.T) {
	fmt.Println("************* TESTIFY *****************")
}

func TestAddWithTestify(t *testing.T) {
	result := Add(1, 2)
	Equal(t, 3, result)
}

func TestSubtractWithTestify(t *testing.T) {
	result := Subtract(5, 3)
	Equal(t, 2, result)
}

func TestMultiplyWithTestify(t *testing.T) {
	result := Multiply(5, 6)
	Equal(t, 30, result)
}

func TestDivideWithTestify(t *testing.T) {
	result := Divide(10, 2)
	Equal(t, float64(5), result)
}
