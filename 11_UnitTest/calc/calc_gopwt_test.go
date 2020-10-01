package calc_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	. "example.com/mittaus/go-testing-frameworks/calc"
	. "github.com/ToQoz/gopwt"
	. "github.com/ToQoz/gopwt/assert"
)

func TestNameGopwt(t *testing.T) {
	fmt.Println("************* GOPWT *****************")
}

func TestMain(m *testing.M) {
	flag.Parse()
	Empower()
	os.Exit(m.Run())
}

func TestAddWithGopwt(t *testing.T) {
	result := Add(1, 2)
	OK(t, 4 == result)
}

func TestSubtractWithGopwt(t *testing.T) {
	result := Subtract(5, 3)
	OK(t, 2 == result)
}

func TestMultiplyWithGopwt(t *testing.T) {
	result := Multiply(5, 6)
	OK(t, 30 == result)
}

func TestDivideWithGopwt(t *testing.T) {
	result := Divide(10, 2)
	OK(t, float64(5) == result)
}
