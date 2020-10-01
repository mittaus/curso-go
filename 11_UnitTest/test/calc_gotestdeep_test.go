package calc_test

import (
	"testing"

	. "example.com/mittaus/go-testing-frameworks/calc"
	. "github.com/maxatome/go-testdeep"
)

func TestAddWithGoTestDeep(t *testing.T) {
	result := Add(1, 2)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(3))
	CmpDeeply(t, result, Code(func(r int) (bool, string) {
		if r == 3 {
			return true, ""
		}
		return false, "Result should be 3"
	}))
}
