package main

import (
	"fmt"
	//"os"
)

//Defer

func primera() {
	fmt.Println("Primera")
}
func segunda() {
	fmt.Println("Segunda")
}

func main() {

	defer primera()
	defer segunda()

	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}

}
